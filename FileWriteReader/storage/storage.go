package storage

import (
	"encoding/json"
	"fmt"
)

type Storage struct {
	Data map[string]interface{} `json:"data"`
}

func New() *Storage {
	return &Storage{
		Data: make(map[string]interface{}),
	}
}

func (s *Storage) Set(key string, value interface{}) {
	s.Data[key] = value
}

func (s *Storage) Get(key string) (interface{}, bool) {
	value, exists := s.Data[key]
	return value, exists
}

func (s *Storage) FromJSON(jsonData []byte) error {
	return json.Unmarshal(jsonData, &s.Data)
}

func (s *Storage) ToJSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

func (s *Storage) String() string {
	return fmt.Sprintf("Storage with %d items", len(s.Data))
}
