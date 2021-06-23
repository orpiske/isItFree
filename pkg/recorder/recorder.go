package recorder

import (
	"log"

	"github.com/orpiske/isItFree/pkg/report"
)

// Recorder interface
type Recorder interface {
	Record(r *report.Report) error
}

// NewRecorder creates a new recorder by adapter
func NewRecorder(adapter string) Recorder {
	switch adapter {
	case "disk":
		return new(DiskRecorder)
	case "influx":
		return new(InfluxRecorder)
	}

	log.Printf("Missing recorder adapter: using default (influxdb)")
	return new(InfluxRecorder)
}
