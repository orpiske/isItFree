package recorder

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/orpiske/isItFree/pkg/report"
	"log"
	"os"
	"time"
)

// InfluxRecorder struct
type InfluxRecorder struct{}

// Record the given data to the recorder
func (i InfluxRecorder) Record(r *report.Report) error {
	bucket := os.Getenv("IIF_BUCKET")
	if len(bucket) == 0 {
		bucket = "academia-test"
	}

	org := os.Getenv("IIF_ORG")
	if len(org) == 0 {
		org = "Home"
	}

	token := os.Getenv("IIF_TOKEN")
	url := os.Getenv("IIF_URL")

	client := influxdb2.NewClient(url, token)
	defer client.Close()

	writeAPI := client.WriteAPIBlocking(org, bucket)

	p := influxdb2.NewPoint("utilization",
		map[string]string{"area": r.Area},
		map[string]interface{}{"used": r.Used, "capacity": r.Capacity},
		time.Now())

	log.Printf("Writing %s data to the DB", r.Area)
	if error := writeAPI.WritePoint(context.Background(), p); error != nil {
		return error
	}

	return nil
}
