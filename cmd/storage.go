package cmd

import (
	"encoding/json"
	"os"
)

func load[T any](stToLoad *T) error {
	file, err := os.ReadFile("./storage.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(file, stToLoad)
}

func save[T any](stToSave T) error {

	file, err := json.MarshalIndent(stToSave, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("./storage.json", file, 0644)
}
