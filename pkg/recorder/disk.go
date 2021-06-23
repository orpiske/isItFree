package recorder

import (
	"encoding/json"
	"fmt"
	"github.com/orpiske/isItFree/pkg/report"
	"io/ioutil"
	"log"
	"os"
)

// DiskRecorder struct
type DiskRecorder struct{}

// Record the given data to the recorder
func (d DiskRecorder) Record(r *report.Report) error {
	path := fmt.Sprintf("%s/isitfree/%s.json", os.TempDir(), r.Area)

	log.Printf("Saving data to %s", path)

	file, _ := json.MarshalIndent(r, "", " ")
	if err := ioutil.WriteFile(path, file, 0644); err != nil {
		log.Printf("Unable to record data to the file.")

		return err
	}

	return nil
}
