package storage

import (
	"encoding/json"
	"os"
)

func LoadIndex() (Index, error) {
	var index Index

	file, err := os.Open("storage/index.json")
	if err != nil {
		return index, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&index)
	return index, err
}
