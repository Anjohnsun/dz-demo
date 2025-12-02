package bins

import (
	"errors"
)

var (
	ErrFileNotFound = errors.New("bin file not found")
)

type BinFile struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Size    int64  `json:"size"`
	Version string `json:"version"`
}
