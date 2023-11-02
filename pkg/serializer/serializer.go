package serializer

type Format string

const (
	FormatJSON Format = "json"
	FormatYAML Format = "yaml"
)

// JSON holds a method to convert an underlying implementation to JSON format.
type JSON interface {
	JSON() ([]byte, error)
	FromJSON(data []byte) error
}

// YAML holds a method to convert an underlying implementation to YAML format.
type YAML interface {
	YAML() ([]byte, error)
	FromYAML(data []byte) error
}

// API holds a method to convert an underlying implementation to its API counterpart.
// This method uses generics to support different API entities.
type API[T any] interface {
	API() T
	FromAPI(in T)
}
