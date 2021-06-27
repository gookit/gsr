package gsr

// Marshaler interface for Marshal/Unmarshal data
type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, ptr interface{}) error
}
