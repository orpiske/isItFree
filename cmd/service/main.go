package main

import (
	"log"
	"os"
	"time"

	"github.com/orpiske/isItFree/pkg/recorder"

	"github.com/orpiske/isItFree/pkg/parser"
	"github.com/orpiske/isItFree/pkg/reader"
)

func main() {
	if len(os.Args) > 1 {
		data, err := reader.FromLocal(os.Args[1])
		if err == nil {
			parser.ParseGym(data)
		}
		return
	}

	gymURL := os.Getenv("IIF_GYM_SOURCE_URL")
	poolURL := os.Getenv("IIF_POOL_SOURCE_URL")

	adapter := os.Getenv("IIF_RECORDER_ADAPTER")
	rec := recorder.NewRecorder(adapter)

	for {

		// @TODO Can be done inside goroutines
		gdata, err := reader.FromWeb(gymURL)
		if err == nil {
			gr, _ := parser.ParseGym(gdata)
			rec.Record(gr)
		}

		// @TODO Can be done inside goroutines
		pdata, err := reader.FromWeb(poolURL)
		if err == nil {
			pr, _ := parser.ParsePool(pdata)
			rec.Record(pr)
		}

		// @FIXME Using sleep is not a good idea
		log.Print("Sleeping for 10 minutes")
		time.Sleep(10 * time.Minute)
	}
}
