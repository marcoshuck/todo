package serializer

// JsonSerializer holds a method to convert an underlying implementation to JSON format.
type JsonSerializer interface {
	JSON() ([]byte, error)
}

// YamlSerializer holds a method to convert an underlying implementation to YAML format.
type YamlSerializer interface {
	YAML() ([]byte, error)
}

// ApiSerializer holds a method to convert an underlying implementation to its API counterpart.
// This method uses generics to support different API entities.
type ApiSerializer[T any] interface {
	API() T
}
