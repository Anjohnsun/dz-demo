package bins

import (
	"FWriteRead/FileWriteReader/file"
	"FWriteRead/FileWriteReader/storage"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrFileNotFound = errors.New("bin file not found")
	ErrInvalidJSON  = errors.New("invalid JSON format")
)

type BinFile struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Size     int64  `json:"size"`
	Checksum string `json:"checksum,omitempty"`
	Version  string `json:"version"`
}

type BinManager struct {
	storage *storage.Storage
}

func NewBinManager() *BinManager {
	return &BinManager{
		storage: storage.New(),
	}
}

func NewBinManagerFromFile(filename string) (*BinManager, error) {
	s, err := storage.NewFromFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create bin manager from file: %w", err)
	}

	return &BinManager{
		storage: s,
	}, nil
}

func (bm *BinManager) LoadBinInfo(filename string) error {
	return bm.storage.LoadFromFile(filename)
}

func (bm *BinManager) SaveBinInfo(filename string) error {
	return bm.storage.SaveToFile(filename)
}

func (bm *BinManager) AddBinFile(binFile BinFile) {
	bm.storage.Set(binFile.Name, binFile)
}

func (bm *BinManager) GetBinFile(name string) (*BinFile, error) {
	value, exists := bm.storage.Get(name)
	if !exists {
		return nil, ErrFileNotFound
	}

	jsonData, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal bin file data: %w", err)
	}

	var binFile BinFile
	if err := json.Unmarshal(jsonData, &binFile); err != nil {
		return nil, fmt.Errorf("failed to unmarshal bin file data: %w", err)
	}

	return &binFile, nil
}

func (bm *BinManager) GetAllBinFiles() ([]BinFile, error) {
	jsonData, err := bm.storage.ToJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to get storage JSON: %w", err)
	}

	var data map[string]BinFile
	if err := json.Unmarshal(jsonData, &data); err != nil {
		var dataList []BinFile
		if err := json.Unmarshal(jsonData, &dataList); err != nil {
			return nil, fmt.Errorf("failed to unmarshal bin files data: %w", err)
		}
		return dataList, nil
	}

	result := make([]BinFile, 0, len(data))
	for _, binFile := range data {
		result = append(result, binFile)
	}

	return result, nil
}

func (bm *BinManager) RemoveBinFile(name string) {
	bm.storage.Delete(name)
}

func (bm *BinManager) SyncWithFileSystem(binDir string) error {
	entries, err := os.ReadDir(binDir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", binDir, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if filepath.Ext(entry.Name()) == ".bin" {
			info, err := entry.Info()
			if err != nil {
				continue
			}

			binFile := BinFile{
				Name:    entry.Name(),
				Path:    filepath.Join(binDir, entry.Name()),
				Size:    info.Size(),
				Version: "1.0.0",
			}

			bm.AddBinFile(binFile)
		}
	}

	return nil
}

func (bm *BinManager) ImportFromJSONFile(filename string) error {
	if !file.IsJSONFile(filename) {
		return fmt.Errorf("file %s is not a JSON file", filename)
	}

	jsonData, err := file.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %w", err)
	}

	if !json.Valid(jsonData) {
		return ErrInvalidJSON
	}

	tempStorage := storage.New()
	if err := tempStorage.FromJSON(jsonData); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	for _, key := range tempStorage.Keys() {
		if value, exists := tempStorage.Get(key); exists {
			bm.storage.Set(key, value)
		}
	}

	return nil
}
