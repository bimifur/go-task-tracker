package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/internal/models"
	"task-tracker/internal/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: task-cli [command] [arguments]")
		fmt.Println("Команды: add, list")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Ошибка! Введите описание задачи")
			return
		}

		description := os.Args[2]
		err := service.AddTask(description)
		if err != nil {
			fmt.Println("Ошибка: ", err)
			return
		} else {
			fmt.Println("Задача успешно добавлена!")
		}

	case "list":
		tasks, err := service.ListTasks()
		if err != nil {
			fmt.Println("Ошибка: ", err)
			return
		}

		if len(os.Args) == 2 {
			for _, t := range tasks {
				fmt.Printf("[ID: %v] %v (status: %v)\n", t.ID, t.Description, t.Status)
			}
		} else if len(os.Args) == 3 {
			status := os.Args[2]
			if !models.IsValidStatus(models.TaskStatus(status)) {
				fmt.Println("Некорректный статус")
				fmt.Printf("Возможные статусы:\ntodo\nin-progress\ndone")
				return
			}

			tasks, err = service.WithStatus(models.TaskStatus(status))
			if len(tasks) == 0 {
				fmt.Printf("Список задач со статусом %v пуст", status)
				return
			}
			if err != nil {
				fmt.Println("Ошибка: ", err)
				return
			}

			for _, t := range tasks {
				fmt.Printf("[ID: %v] %v (status: %v)\n", t.ID, t.Description, t.Status)
			}
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Ошибка! Введите ID и статус!")
			return
		}

		ID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Ошибка: ", err)
			return
		}

		statusFromArgs := os.Args[3]
		status := models.TaskStatus(statusFromArgs)

		if !models.IsValidStatus(status) {
			fmt.Printf("invalid status: %s. Use todo, in-progress or done", status)
			return
		}

		err = service.UpdateTask(ID, status)
		if err != nil {
			fmt.Println("Ошибка: ", err)
			return
		}
		fmt.Printf("Задача успешно обновлена! (ID: %v)", ID)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Ошибка! Введите ID")
			return
		}

		intID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Ошибка: ", err)
			return
		}

		err = service.DeleteTask(intID)
		fmt.Printf("Задача с ID %v удалена", intID)

	default:
		fmt.Println("Неизвестная команда")
		return
	}
}
