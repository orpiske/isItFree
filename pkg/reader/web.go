package reader

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// FromWeb read the data from web resource
func FromWeb(name string, url string) ([]byte, error) {
	if len(url) == 0 {
		return nil, errors.New(fmt.Sprintf("The %s source URL is unset", name))
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Unexpected status from the server: %d %s", res.StatusCode, res.Status))
	}

	return ioutil.ReadAll(res.Body)
}
