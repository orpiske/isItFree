package recorder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/orpiske/isItFree/pkg/report"
)

// DiskRecorder struct
type DiskRecorder struct{}

// Record the given data to the recorder
func (d DiskRecorder) Record(r *report.Report) {
	path := fmt.Sprintf("/tmp/isitfree/%s.json", r.Area)

	log.Printf("Saving data to %s", path)

	file, _ := json.MarshalIndent(r, "", " ")
	if err := ioutil.WriteFile(path, file, 0644); err != nil {
		log.Printf("Unable to record data to the file.")
	}
}
