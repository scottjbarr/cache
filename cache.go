package cache

// Reader is an interface limited to reading a Cache.
type Reader interface {
	Get(string) (interface{}, error)
}

// Writer is an interface limited to writing to a Cache.
type Writer interface {
	Set(string, interface{}) error
}

// Cache defines the interface for reading an writing to a Cache.
type Cache interface {
	Reader
	Writer
}
