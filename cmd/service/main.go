package main

import (
	"log"
	"os"
	"time"

	"github.com/orpiske/isItFree/pkg/parser"
	"github.com/orpiske/isItFree/pkg/reader"
)

func main() {
	if len(os.Args) > 1 {
		data, err := reader.FromLocal(os.Args[1])
		if err != nil {
			return
		}
		parser.ParseGym(data)
		return
	}

	gymURL := os.Getenv("IIF_GYM_SOURCE_URL")
	poolURL := os.Getenv("IIF_POOL_SOURCE_URL")

	for true {

		// @TODO Can be done inside goroutines
		gdata, err := reader.FromWeb(gymURL)
		if err != nil {
			return
		}
		parser.ParseGym(gdata)

		// @TODO Can be done inside goroutines
		pdata, err := reader.FromWeb(poolURL)
		if err != nil {
			return
		}
		parser.ParsePool(pdata)

		// @FIXME Using sleep is not a good idea
		log.Print("Sleeping for 10 minutes")
		time.Sleep(10 * time.Minute)
	}
}
