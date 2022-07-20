package gsr

// Marshaler interface
type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
}

// Unmarshaler interface
type Unmarshaler interface {
	Unmarshal(v []byte, ptr interface{}) error
}

// DataParser interface for Marshal/Unmarshal data
type DataParser interface {
	Marshaler
	Unmarshaler
}

// MarshalFunc define
type MarshalFunc func(v interface{}) ([]byte, error)

// Marshal implements the Marshaler
func (m MarshalFunc) Marshal(v interface{}) ([]byte, error) {
	return m(v)
}

// UnmarshalFunc define
type UnmarshalFunc func(v []byte, ptr interface{}) error

// Unmarshal implements the Unmarshaler
func (u UnmarshalFunc) Unmarshal(v []byte, ptr interface{}) error {
	return u(v, ptr)
}
