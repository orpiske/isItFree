package recorder

import (
	"github.com/orpiske/isItFree/pkg/report"
)

// Recorder interface
type Recorder interface {
	Record(r *report.Report)
}
