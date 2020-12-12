package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func parsePool(r io.Reader) {
	log.Print("Trying to find pool utilization")

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Print("Unable to parse the document: " + err.Error())
		return
	}

	var current Capacity

	current.area = "Kravy Hora (interna)"
	current.capacity = 165
	log.Printf("Area: %s\n", current.area)
	doc.Find(".field-items").Each(func(i int, s *goquery.Selection) {
		s.Find("p:nth-child(7)").Find("strong").Each(func(i int, s *goquery.Selection) {

			fullTxt := strings.TrimSpace(s.Text())
			current.used, _ = strconv.ParseInt(fullTxt, 10, 8)
		})
	})

	log.Printf("Used: %d\n", current.used)
	log.Printf("Capacity: %d\n", current.capacity)
	record(current)
}

func readPoolFromWeb() {
	targetUrl := os.Getenv("IIF_POOL_SOURCE_URL")

	if len(targetUrl) == 0 {
		log.Printf("The source URL for the pool is unset")
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

	parsePool(res.Body)
}
