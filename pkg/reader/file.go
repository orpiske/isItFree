package reader

import (
	"bufio"
	"log"
	"os"
)

// FromLocal read data from file resource
func FromLocal(fn string) ([]byte, error) {
	log.Print("Reading from file")

	f, err := os.Open(fn)
	if err != nil {
		log.Printf("Unable to open file: %s\n", err.Error())
		return nil, err
	}

	defer cleanup(f)
	r := bufio.NewScanner(f)

	return r.Bytes(), nil
}

func cleanup(f *os.File) {
	if f != nil {
		err := f.Close()

		if err != nil {
			log.Printf("Error closing file: %s\n", err.Error())
		}
	}
}
