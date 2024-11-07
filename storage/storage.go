package storage

// Dependency Inversion

type Storage interface {
	Read() (string, error)
	Write(data string) error
}
