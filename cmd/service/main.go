package main

import (
	"github.com/orpiske/isItFree/pkg/parser"
	"github.com/orpiske/isItFree/pkg/reader"
	"github.com/orpiske/isItFree/pkg/recorder"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	// Collect one time at launch to avoid waiting for the ticker
	collect(rec, gymURL, poolURL)

	// Run subsequent collections using the ticker
	scheduledCollection(rec, gymURL, poolURL)
}

func scheduledCollection(rec recorder.Recorder, gymURL string, poolURL string) {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-signals:
				log.Print("Terminating execution because of signals")
				done <- true
				return
			case <-ticker.C:
				collect(rec, gymURL, poolURL)
				log.Print("Waiting for the next update ...")
			}
		}
	}()

	<-done
}

func collect(rec recorder.Recorder, gymURL string, poolURL string) {

	if gdata, readerErr := reader.FromWeb("gym", gymURL) ; readerErr != nil {
		log.Print(readerErr.Error())
	} else {
		log.Print("Parsing gym data")
		if gr, parserErr := parser.ParseGym(gdata) ; parserErr != nil {
			log.Print(readerErr.Error())
		} else {
			log.Print("Recording gym data")
			if recorderErr := rec.Record(gr) ; recorderErr != nil {
				log.Printf(recorderErr.Error())
			}
		}
	}

	if pdata, readerErr := reader.FromWeb("pool", poolURL) ; readerErr != nil {
		log.Print(readerErr.Error())
	} else {
		if pr, parserErr := parser.ParsePool(pdata) ; parserErr != nil {
			log.Print(readerErr.Error())
		} else {
			log.Print("Recording pool data")
			if recorderErr := rec.Record(pr) ; recorderErr != nil {
				log.Printf(recorderErr.Error())
			}
		}
	}
}
