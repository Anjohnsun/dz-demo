package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

func WriteFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func AppendFile(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

func IsJSONFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".json"
}

func EnsureJSONExtension(filename string) string {
	if IsJSONFile(filename) {
		return filename
	}
	return filename + ".json"
}

func ValidateJSONFile(filename string) error {
	if !Exists(filename) {
		return os.ErrNotExist
	}
	if !IsJSONFile(filename) {
		return &InvalidExtensionError{Filename: filename, Expected: ".json"}
	}
	return nil
}

type InvalidExtensionError struct {
	Filename string
	Expected string
}

func (e *InvalidExtensionError) Error() string {
	ext := filepath.Ext(e.Filename)
	return fmt.Sprintf("file %s has extension %s, expected %s",
		e.Filename, ext, e.Expected)
}
