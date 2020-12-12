package main

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func parseGym(r io.Reader) {
	log.Print("Trying to find gym utilization")

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Print("Unable to parse the document: " + err.Error())
		return
	}

	var current Capacity

	doc.Find(".s-box").Each(func(i int, s *goquery.Selection) {

		s.Find(".s-box__body__content").Find("h3").Each(func(i int, s *goquery.Selection) {
			current.area = strings.TrimSpace(s.Text())

			log.Printf("Area: %s\n", current.area)
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

		log.Printf("Used: %d\n", current.used)
		log.Printf("Capacity: %d\n", current.capacity)

		record(current)
	})
}

func readGymFromWeb() {
	targetUrl := os.Getenv("IIF_GYM_SOURCE_URL")

	if len(targetUrl) == 0 {
		log.Printf("The source URL for the gym is unset")
		return
	}

	res, err := http.Get(targetUrl)
	if err != nil {
		log.Print("The connection might be down: " + err.Error())
		log.Print("Skipping this update and resuming on the next")
		return
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
		log.Print("Skipping this update and resuming on the next")
		return
	}

	parseGym(res.Body)
}

func cleanup(f *os.File) {
	if f != nil {
		err := f.Close()

		if err != nil {
			log.Printf("Error closing file: %s\n", err.Error())
		}
	}
}

func readGymFromLocal() {
	log.Print("Reading from file")
	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Printf("Unable to open file: %s\n", err.Error())
		return
	}

	defer cleanup(f)

	r := bufio.NewReader(f)

	parseGym(r)
}
