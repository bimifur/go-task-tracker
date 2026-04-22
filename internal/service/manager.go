package service

import (
	"fmt"
	"task-tracker/internal/models"
	"task-tracker/internal/storage"
	"time"
)

func AddTask(description string) error {
	tasks, err := storage.Load()
	if err != nil {
		return err
	}

	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	newID := maxID + 1

	newTask := models.Task{
		ID:          newID,
		Description: description,
		Status:      models.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	return storage.Save(tasks)
}

func ListTasks() ([]models.Task, error) {
	tasks, err := storage.Load()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func UpdateTask(ID int, newStatus models.TaskStatus) error {
	tasks, err := storage.Load()
	if err != nil {
		return err
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == ID {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if found == false {
		return fmt.Errorf("task with id %d not found", ID)
	}

	return storage.Save(tasks)
}

func DeleteTask(ID int) error {
	tasks, err := storage.Load()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return storage.Save(tasks)
		}
	}
	return fmt.Errorf("нет такого ID: %v", ID)
}
