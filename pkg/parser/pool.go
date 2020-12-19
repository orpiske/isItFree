package parser

import (
	"bytes"
	"log"
	"strconv"
	"strings"

	"github.com/orpiske/isItFree/pkg/report"

	"github.com/PuerkitoBio/goquery"
)

// ParsePool data
func ParsePool(b []byte) (*report.Report, error) {
	log.Print("Trying to find pool utilization")

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		log.Print("Unable to parse the document: " + err.Error())
		return nil, err
	}

	report := &report.Report{}

	report.Area = "Kravy Hora (interna)"
	report.Capacity = 165
	log.Printf("Area: %s\n", report.Area)
	doc.Find(".field-items").Each(func(i int, s *goquery.Selection) {
		s.Find("p:nth-child(7)").Find("strong").Each(func(i int, s *goquery.Selection) {

			fullTxt := strings.TrimSpace(s.Text())
			report.Used, _ = strconv.ParseInt(fullTxt, 10, 8)
		})
	})

	log.Printf("Used: %d\n", report.Used)
	log.Printf("Capacity: %d\n", report.Capacity)

	return report, nil
}
