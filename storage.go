package main

import (
	"encoding/json"
	"os"
)

func loadTasks() ([]Task, error) {
	//Check if the file exists
	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
		return []Task{}, nil
	}

	//Read the file
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return nil, err
	}

	//Handle empty file
	if len(data) == 0 {
		return []Task{}, nil
	}

	//Parse the file
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTasks(tasks []Task) error {
	//Convert structs to JSON
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	//Overwrite whole file with all tasks data
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}

	return nil //Successfully saved tasks
}
