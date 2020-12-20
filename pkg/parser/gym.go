package parser

import (
	"bytes"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/orpiske/isItFree/pkg/report"
)

// ParseGym data
func ParseGym(b []byte) (*report.Report, error) {
	log.Print("Trying to find gym utilization")

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		log.Print("Unable to parse the document: " + err.Error())
		return nil, err
	}

	report := &report.Report{}

	doc.Find(".s-box").Each(func(i int, s *goquery.Selection) {

		s.Find(".s-box__body__content").Find("h3").Each(func(i int, s *goquery.Selection) {
			report.Area = strings.TrimSpace(s.Text())

			log.Printf("Area: %s\n", report.Area)
		})

		s.Find("span").Each(func(i int, s *goquery.Selection) {
			txtCapacityMax := strings.TrimSpace(s.Text())

			if !strings.Contains(txtCapacityMax, "/") {
				return
			}

			txtUsed := strings.TrimSpace(strings.Split(txtCapacityMax, "/")[0])
			txtCapacity := strings.TrimSpace(strings.Split(txtCapacityMax, "/")[1])

			report.Used, _ = strconv.ParseInt(txtUsed, 10, 8)
			report.Capacity, _ = strconv.ParseInt(txtCapacity, 10, 8)
		})

		log.Printf("Used: %d\n", report.Used)
		log.Printf("Capacity: %d\n", report.Capacity)
	})

	return report, nil
}
