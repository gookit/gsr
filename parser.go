package gsr


// Serializer interface definition
type Serializer interface {
	Serialize(v any) ([]byte, error)
	Deserialize(data []byte, v any) error
}

// GoSerializer interface definition
type GoSerializer interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(v []byte, ptr any) error
}

// DataParser interface
// Deperacted
type DataParser = GoSerializer

// Codec interface definition
type Codec interface {
	Decode(blob []byte, v any) (err error)
	Encode(v any) (out []byte, err error)
}

//
//
//

// Marshaler interface
type Marshaler interface {
	Marshal(v any) ([]byte, error)
}

// Unmarshaler interface
type Unmarshaler interface {
	Unmarshal(v []byte, ptr any) error
}

// MarshalFunc define
type MarshalFunc func(v any) ([]byte, error)

// Marshal implements the Marshaler
func (m MarshalFunc) Marshal(v any) ([]byte, error) {
	return m(v)
}

// UnmarshalFunc define
type UnmarshalFunc func(v []byte, ptr any) error

// Unmarshal implements the Unmarshaler
func (u UnmarshalFunc) Unmarshal(v []byte, ptr any) error {
	return u(v, ptr)
}

// DecodeFunc definition
type DecodeFunc func(blob []byte, v any) (err error)

// EncodeFunc definition
type EncodeFunc func(v any) (out []byte, err error)

