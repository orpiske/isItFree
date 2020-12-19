package recorder

import (
	"context"
	"os"
	"time"
)

// InfluxRecorder struct
type InfluxRecorder struct{}

// Record the given data to the recorder
func (r *InfluxRecorder) Record(current Capacity) error {
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
	writeAPI := client.WriteAPIBlocking(org, bucket)

	p := influxdb2.NewPoint("utilization",
		map[string]string{"area": current.area},
		map[string]interface{}{"used": current.used, "capacity": current.capacity},
		time.Now())

	writeAPI.WritePoint(context.Background(), p)

	client.Close()

	return nil
}
