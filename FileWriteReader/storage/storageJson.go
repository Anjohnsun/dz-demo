package storage

import (
	"FWriteRead/FileWriteReader/file"
	"encoding/json"
	"fmt"
)

type StorageJson struct {
	Data map[string]interface{} `json:"data"`
}

func NewJson() *StorageJson {
	return &StorageJson{
		Data: make(map[string]interface{}),
	}
}

func (s *StorageJson) Set(key string, value interface{}) {
	s.Data[key] = value
}

func (s *StorageJson) Get(key string) (interface{}, bool) {
	value, exists := s.Data[key]
	return value, exists
}

func (s *StorageJson) Delete(key string) {
	delete(s.Data, key)
}

func (s *StorageJson) Load(filename string) error {
	if !file.IsJSONFile(filename) {
		return fmt.Errorf("file %s is not a JSON file", filename)
	}

	if !file.Exists(filename) {
		return fmt.Errorf("file %s does not exist", filename)
	}

	jsonData, err := file.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if !json.Valid(jsonData) {
		return fmt.Errorf("invalid JSON format")
	}

	return json.Unmarshal(jsonData, &s.Data)
}

func (s *StorageJson) Save(filename string) error {
	filename = file.EnsureJSONExtension(filename)

	jsonData, err := json.MarshalIndent(s.Data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	if err := file.WriteFile(filename, jsonData); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
