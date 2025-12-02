package storage

import (
	"FWriteRead/FileWriteReader/file"
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

func NewFromFile(filename string) (*Storage, error) {
	s := New()
	if err := s.LoadFromFile(filename); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Storage) Set(key string, value interface{}) {
	s.Data[key] = value
}

func (s *Storage) Get(key string) (interface{}, bool) {
	value, exists := s.Data[key]
	return value, exists
}

func (s *Storage) Delete(key string) {
	delete(s.Data, key)
}

func (s *Storage) Keys() []string {
	keys := make([]string, 0, len(s.Data))
	for k := range s.Data {
		keys = append(keys, k)
	}
	return keys
}

func (s *Storage) Clear() {
	s.Data = make(map[string]interface{})
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

func (s *Storage) SaveToFile(filename string) error {
	filename = file.EnsureJSONExtension(filename)

	jsonData, err := s.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal storage to JSON: %w", err)
	}

	if err := file.WriteFile(filename, jsonData); err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filename, err)
	}

	return nil
}

func (s *Storage) LoadFromFile(filename string) error {
	if !file.IsJSONFile(filename) {
		return fmt.Errorf("file %s is not a JSON file (.json extension required)", filename)
	}

	if !file.Exists(filename) {
		return fmt.Errorf("file %s does not exist", filename)
	}

	jsonData, err := file.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	if !json.Valid(jsonData) {
		return fmt.Errorf("invalid JSON format in file %s", filename)
	}

	if err := s.FromJSON(jsonData); err != nil {
		return fmt.Errorf("failed to unmarshal JSON from file %s: %w", filename, err)
	}

	return nil
}

func (s *Storage) SaveToFileWithValidation(filename string) error {
	if err := file.ValidateJSONFile(filename); err == nil {
	}

	return s.SaveToFile(filename)
}
