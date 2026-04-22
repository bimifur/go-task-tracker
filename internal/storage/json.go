package storage

import (
	"encoding/json"
	"os"
	"task-tracker/internal/models"
)

func Load() ([]models.Task, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		} else {
			return nil, err
		}
	}

	if len(data) == 0 {
		return []models.Task{}, nil
	}

	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func Save(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
