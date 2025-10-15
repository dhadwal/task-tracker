package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Print("Command missing\nUsage: task-tracker <command>\n")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Print("Description missing\nUsage: task-tracker add <description>\n")
			return
		}
		description := os.Args[2]
		fmt.Printf("Adding new task: %s\n", description)

		err := addTask(description)
		if err != nil {
			fmt.Printf("Error adding task: %s\n", err)
			return
		}

	case "update":
		if len(os.Args) < 3 {
			fmt.Print("ID missing\nUsage: task-tracker update <id> <description>\n")
			return
		}
		if len(os.Args) < 4 {
			fmt.Print("Description missing\nUsage: task-tracker update <id> <description>\n")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Invalid ID: %s\n", os.Args[2])
			return
		}
		description := os.Args[3]
		fmt.Printf("Updating task with ID %d to %s\n", id, description)
		err = updateTask(id, description)
		if err != nil {
			fmt.Printf("Error updating task: %s\n", err)
			return
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Print("ID missing\nUsage: task-tracker delete <id>\n")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Invalid ID: %s\n", os.Args[2])
			return
		}
		fmt.Printf("Deleting task with ID: %d\n", id)
		err = deleteTask(id)
		if err != nil {
			fmt.Printf("Error deleting task: %s\n", err)
			return
		}
	case "list":
		status := ""
		if len(os.Args) == 3 {
			status = os.Args[2]
			fmt.Printf("Listing tasks with status: %s\n", status)
		} else {
			fmt.Printf("Listing tasks...\n")
		}
		err := listTasks(status)
		if err != nil {
			fmt.Printf("Error listing tasks: %s\n", err)
			return
		}
	case "mark":
		if len(os.Args) < 3 {
			fmt.Printf("Status missing\nUsage: task-tracker mark <status> <id>\n")
			return
		}
		if len(os.Args) < 4 {
			fmt.Printf("ID missing\nUsage: task-tracker mark <status> <id>\n")
			return
		}
		status := os.Args[2]
		if status != "todo" && status != "in-progress" && status != "done" {
			fmt.Printf("Invalid status: %s\n", status)
			return
		}
		id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("Invalid ID: %s\n", os.Args[3])
			return
		}
		fmt.Printf("Marking task with ID %d as %s\n", id, status)
		err = markTask(status, id)
		if err != nil {
			fmt.Printf("Error marking task: %s\n", err)
			return
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}

}
