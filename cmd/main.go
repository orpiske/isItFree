package main

import (
	"bufio"
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Capacity struct {
	collectionTime int64
	area string
	used int64
	capacity int64
}

func record(current Capacity) {
	bucket := os.Getenv("IIF_BUCKET")
	org := os.Getenv("IIF_ORG")
	token := os.Getenv("IIF_TOKEN")
	url := os.Getenv("IIF_URL")

	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPIBlocking(org, bucket)

	p := influxdb2.NewPoint("stat",

		map[string]string{"area": current.area},
		map[string]interface{}{"used": current.used, "capacity": current.capacity},
		time.Now())

	writeAPI.WritePoint(context.Background(), p)

	client.Close()
}

func parse(r io.Reader) {
	// Load the HTML document
	//doc, err := goquery.NewDocumentFromReader(res.Body)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}

	var current Capacity;

	// Find the review items s-col s-col--4
	doc.Find(".s-box").Each(func(i int, s *goquery.Selection) {

		s.Find(".s-box__body__content").Find("h3").Each(func(i int, s *goquery.Selection) {
			current.area = strings.TrimSpace(s.Text())

			fmt.Printf("Area: %s\n", current.area)
		})

		s.Find("span").Each(func(i int, s *goquery.Selection) {
			txtCapacityMax := strings.TrimSpace(s.Text())

			if !strings.Contains(txtCapacityMax, "/") {
				return
			}

			txtUsed := strings.TrimSpace(strings.Split(txtCapacityMax, "/")[0])
			txtCapacity := strings.TrimSpace(strings.Split(txtCapacityMax, "/")[1])

			current.used, _ = strconv.ParseInt(txtUsed, 10, 8)
			current.capacity, _ = strconv.ParseInt(txtCapacity, 10, 8)
		})

		fmt.Printf("Used: %d\n", current.used)
		fmt.Printf("Capacity: %d\n", current.capacity)

		record(current)
	})
}

func main() {


	if len(os.Args) > 1 {
		readFromLocal()

	} else {
		for true {
			readFromWeb()
			fmt.Printf("Sleeping for 10 minutes")
			time.Sleep(10 * time.Minute)
		}

	}
}

func readFromWeb() {
	targetUrl := os.Getenv("IIF_SOURCE_URL")

	if len(targetUrl) == 0 {
		targetUrl = "https://rasinova.sportujemevbrne.cz/"
	}

	res, err := http.Get(targetUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	parse(res.Body)
}

func readFromLocal() {
	fmt.Printf("Reading from file\n")
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("Unable to open file %s", os.Args[1])
	}

	defer f.Close()

	r := bufio.NewReader(f)

	parse(r)
}
