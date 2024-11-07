package note

import (
	"coba/storage"
	"fmt"
)

// Read harus mengembalikan string dan (error atau nil)
// Write harus mengembalikan error atau nil
// Semua string yang di "write" harus ada "prefix" nya
// Semua string yang di "write" harus bisa muncul saat "read" dijalankan

type Note struct {
	Storage storage.Storage
	Prefix  string
}

func (n Note) ReadNote() (string, error) {
	return n.Storage.Read()
}

func (n Note) WriteNote(data string) error {
	return n.Storage.Write(fmt.Sprintf("%s %s\n", n.Prefix, data))
}
