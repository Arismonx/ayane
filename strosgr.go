package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Generics
type Storage[T any] struct {
	FileName string
}

// Constructor
func NewStroage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

// functione Save data
func (s *Storage[T]) Save(data *T) error {
	jsonData, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return fmt.Errorf("unable to convert data to JSON: %w", err)
	}

	// if err = os.WriteFile(s.FileName, jsonData, 0644); err != nil {
	// 	return fmt.Errorf("unable to write file '%s' :%w", s.FileName, err)
	// }

	// fmt.Printf("save data in file '%s' complete!", s.FileName)
	// return nil
	return os.WriteFile(s.FileName, jsonData, 0644)

}

// functione Load data
func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)

	if err != nil {
		return fmt.Errorf("unable to read file '%s' :%w", s.FileName, err)
	}

	return json.Unmarshal(fileData, data)
}
