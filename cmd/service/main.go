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

	ticker := time.NewTicker(10 * time.Second)
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
				collect(gymURL, rec, poolURL)
				log.Print("Waiting for the next update ...")
			}
		}
	}()

	<-done
}

func collect(gymURL string, rec recorder.Recorder, poolURL string) {
	gdata, err := reader.FromWeb("gym", gymURL)
	if err == nil {
		gr, _ := parser.ParseGym(gdata)
		rec.Record(gr)
	} else {
		log.Print(err.Error())
	}

	pdata, err := reader.FromWeb("pool", poolURL)
	if err == nil {
		pr, _ := parser.ParsePool(pdata)
		rec.Record(pr)
	} else {
		log.Print(err.Error())
	}
}
