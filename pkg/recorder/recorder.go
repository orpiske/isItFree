package recorder

import (
	"log"

	"github.com/orpiske/isItFree/pkg/report"
)

// Recorder interface
type Recorder interface {
	Record(r *report.Report)
}

// NewRecorder creates a new recorder by adapter
func NewRecorder(adapter string) Recorder {
	if len(adapter) == 0 {
		return new(InfluxRecorder)
	}

	switch adapter {
	case "influx":
		return new(InfluxRecorder)
	case "disk":
		return new(DiskRecorder)
	default:
		log.Printf("Missing recorder adapter.")
		return nil
	}
}
