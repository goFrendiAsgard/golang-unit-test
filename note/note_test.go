package note

import (
	"strings"
	"testing"
)

type MockStorage struct {
	data      []string
	writeCall []string
}

func (m *MockStorage) Read() (string, error) {
	return strings.Join(m.data, ""), nil
}

func (m *MockStorage) Write(data string) error {
	m.writeCall = append(m.writeCall, data)
	m.data = append(m.data, data)
	return nil
}

func TestNote(t *testing.T) {
	s := &MockStorage{
		data:      []string{},
		writeCall: []string{},
	}
	n := Note{Storage: s, Prefix: "abc"}
	n.WriteNote("Kun ke warung kuning")
	// saat n.WriteNote dipanggil, maka s.Write harus terpanggil juga
	if len(s.writeCall) != 1 {
		t.Errorf("s.Write Tidak terpanggil")
	}
	// saat n.WriteNote dipanggil, maka s.Write harus dipanggil dengan parameter prefix + data asli
	if s.writeCall[0] != "abc Kun ke warung kuning\n" {
		t.Errorf("Incorrect s.Write call parameter: %#v", s.writeCall)
	}
	// saat n.Read() dipanggil, maka semua data yang pernah di write harus muncul
	n.WriteNote("disuruh Pak Abdul")
	data, err := n.ReadNote()
	if err != nil {
		t.Errorf("%s", err)
	}
	if data != "abc Kun ke warung kuning\nabc disuruh Pak Abdul\n" {
		t.Errorf("Unexpected result: %s", data)
	}
}
