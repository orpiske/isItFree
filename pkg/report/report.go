package report

// Report struct
type Report struct {
	CollectionTime int64  `json:"collection_time"`
	Area           string `json:"area"`
	Used           int64  `json:"used"`
	Capacity       int64  `json:"capacity"`
}
