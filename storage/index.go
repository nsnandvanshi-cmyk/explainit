package storage

import (
	"encoding/json"
	"os"
)

type Index struct {
	Files []string `json:"files"`
}

func SaveIndex(index Index) error {
	file, err := os.Create("storage/index.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(index)
}
