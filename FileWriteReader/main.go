package main

import (
	"FWriteRead/FileWriteReader/interfaces"
	"FWriteRead/FileWriteReader/storage"
)

func main() {
	var s interfaces.Storage = storage.NewJson()
	s.Load("anyFile.json")
	s.Get("redBird")
	s.Delete("redBird")
	s.Save("anyFile.json")
}
