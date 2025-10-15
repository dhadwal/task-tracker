package main

import (
	"fmt"
	"strconv"
	"time"
)

type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func addTask(description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	//Create new task
	newTask := Task{
		ID:          generateID(tasks),
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	//Add new task to tasks slice
	tasks = append(tasks, newTask)

	//Save tasks to file
	err = saveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("Task added successfully with ID: %s\n", newTask.ID)
	return nil //Successfully added task
}

func generateID(tasks []Task) string {
	maxID := 0
	for _, task := range tasks {
		taskID, err := strconv.Atoi(task.ID)
		if err != nil {
			continue
		}
		if taskID > maxID {
			maxID = taskID
		}
	}
	return strconv.Itoa(maxID + 1)
}

func listTasks(status string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	for _, task := range tasks {
		if status == "" {
			fmt.Printf("ID: %s\nDescription: %s\nStatus: %s\nCreatedAt: %s\nUpdatedAt: %s\n\n", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		} else if status == task.Status {
			fmt.Printf("ID: %s\nDescription: %s\nStatus: %s\nCreatedAt: %s\nUpdatedAt: %s\n\n", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
	return nil
}

func updateTask(id int, description string) error {
	found := false
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == strconv.Itoa(id) {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if found == false {
		return fmt.Errorf("task with ID: %d not found", id)
	}

	err = saveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("Task updated successfully with ID: %d\n", id)
	return nil
}

func deleteTask(id int) error {
	found := false
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == strconv.Itoa(id) {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}

	if found == false {
		return fmt.Errorf("task with ID: %d not found", id)
	}

	err = saveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("Task deleted successfully with ID: %d\n", id)

	return nil
}

func markTask(status string, id int) error {
	found := false
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == strconv.Itoa(id) {
			found = true
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			break
		}
	}

	if found == false {
		return fmt.Errorf("task with ID: %d not found", id)
	}
	err = saveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("Task marked successfully with ID: %d\n", id)
	return nil
}
