package main

import (
	"coba/note"
	"coba/storage"
	"fmt"
)

func main() {
	fs := storage.FileStorage{FileName: "coba.txt"}
	n := note.Note{Storage: fs, Prefix: "Catatan Rekon:"}
	if err := n.WriteNote("Kun ke warung kuning"); err != nil {
		fmt.Println(err)
	}
	if err := n.WriteNote("Diajak Pak Abdul"); err != nil {
		fmt.Println(err)
	}
	data, err := n.ReadNote()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
