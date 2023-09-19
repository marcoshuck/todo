package serializer

// jsonSerializer holds a method to convert an underlying implementation to JSON format.
type jsonSerializer interface {
	JSON() ([]byte, error)
}

// yamlSerializer holds a method to convert an underlying implementation to YAML format.
type yamlSerializer interface {
	YAML() ([]byte, error)
}

// apiSerializer holds a method to convert an underlying implementation to its API counterpart.
// This method uses generics to support different API entities.
type apiSerializer[T any] interface {
	API() T
}
