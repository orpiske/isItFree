package recorder

// Recorder interface
type Recorder interface {
	Record(current Capacity) error
}
