package main

import (
	"log"
	"os"
	"time"
)

type Capacity struct {
	collectionTime int64
	area           string
	used           int64
	capacity       int64
}

func main() {
	if len(os.Args) > 1 {
		readGymFromLocal()

	} else {
		for true {
			readGymFromWeb()
			readPoolFromWeb()
			log.Print("Sleeping for 10 minutes")
			time.Sleep(10 * time.Minute)
		}

	}
}
