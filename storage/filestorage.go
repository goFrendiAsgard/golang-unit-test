package storage

import (
	"os"
)

type FileStorage struct {
	FileName string
}

func (s FileStorage) Read() (string, error) {
	bData, err := os.ReadFile(s.FileName)
	result := string(bData)
	return result, err
}

func (s FileStorage) Write(data string) error {
	// TODO: jangan ditiru wkwkk
	os.OpenFile(s.FileName, os.O_CREATE|os.O_EXCL, 0644)
	file, err := os.OpenFile(s.FileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}
