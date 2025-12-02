package bins

import (
	"FWriteRead/FileWriteReader/storage"
	"errors"
)

var (
	ErrFileNotFound = errors.New("bin file not found")
	ErrInvalidJSON  = errors.New("invalid JSON format")
)

// BinFile представляет информацию о bin файле
type BinFile struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Size     int64  `json:"size"`
	Checksum string `json:"checksum,omitempty"`
	Version  string `json:"version"`
}

// BinManager управляет bin файлами через хранилище
type BinManager struct {
	storage *storage.Storage
}
