package todo

import (
	"encoding/json"
	"errors"
	"os"
)

const filename = "tasks.json"

// ✅ LoadTasks reads tasks from the JSON file
func LoadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil // Return empty list if file doesn't exist
		}
		return nil, err
	}
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

// ✅ SaveTasks writes tasks to the JSON file
func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
