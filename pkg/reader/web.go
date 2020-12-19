package reader

import (
	"errors"
	"io"
	"log"
	"net/http"
)

// FromWeb read the data from web resource
func FromWeb(url string) (io.ReadCloser, error) {
	if len(url) == 0 {
		log.Printf("The source URL is unset")
		return nil, errors.New("The source URL is unset")
	}

	res, err := http.Get(url)
	if err != nil {
		log.Print("The connection might be down: " + err.Error())
		log.Print("Skipping this update and resuming on the next")
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
		log.Print("Skipping this update and resuming on the next")
		return nil, err
	}

	return res.Body, nil
}
