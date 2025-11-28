package main

import (
	"errors"
	"time"
)

type Bin struct {
	Id        string    `json:"id,omitempty"`
	Private   bool      `json:"private,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name,omitempty"`
}

func newBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	if id != "" && name != "" {
		return &Bin{id, private, createdAt, name}, nil
	}
	return nil, errors.New("incorrect id/name")
}

var BinList []Bin

func main() {
	BinList = make([]Bin, 0)
}
