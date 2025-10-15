# Task Tracker CLI

A simple and efficient command-line interface (CLI) application for managing your tasks. Built with Go, this tool helps you keep track of your to-do items with ease.

## Features

- ✅ Add new tasks with descriptions
- 📝 Update existing task descriptions
- 🗑️ Delete tasks
- 📋 List all tasks or filter by status
- 🔄 Mark tasks with different statuses (todo, in-progress, done)
- 💾 Persistent storage using JSON file
- ⏰ Automatic timestamps for creation and updates

## Requirements

- Go 1.25.1 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/dhadwal/task-tracker.git
cd task-tracker
```

2. Build the application:
```bash
go build -o task-tracker
```

3. (Optional) Move to your PATH for global access:
```bash
mv task-tracker /usr/local/bin/
```

## Usage

### Add a Task
Add a new task with a description:
```bash
./task-tracker add "Buy groceries"
```

### List Tasks
List all tasks:
```bash
./task-tracker list
```

List tasks by status:
```bash
./task-tracker list todo
./task-tracker list in-progress
./task-tracker list done
```

### Update a Task
Update a task's description by ID:
```bash
./task-tracker update 1 "Buy groceries and cook dinner"
```

### Mark Task Status
Change the status of a task:
```bash
./task-tracker mark in-progress 1
./task-tracker mark done 1
./task-tracker mark todo 1
```

### Delete a Task
Delete a task by ID:
```bash
./task-tracker delete 1
```

## Project Structure

```
task-tracker/
├── main.go       # Entry point and command routing
├── task.go       # Task operations (add, update, delete, list, mark)
├── storage.go    # File I/O operations for persistent storage
├── go.mod        # Go module file
└── tasks.json    # JSON file storing all tasks (auto-generated)
```

## Task Data Structure

Each task contains:
- `id`: Unique identifier (auto-generated)
- `description`: Task description
- `status`: Current status (todo, in-progress, done)
- `created_at`: Timestamp when task was created
- `updated_at`: Timestamp when task was last modified

## Example Workflow

```bash
# Add some tasks
./task-tracker add "Write project documentation"
./task-tracker add "Review pull requests"
./task-tracker add "Deploy to production"

# List all tasks
./task-tracker list

# Start working on a task
./task-tracker mark in-progress 1

# Update task description
./task-tracker update 1 "Write comprehensive project documentation"

# Complete the task
./task-tracker mark done 1

# List only completed tasks
./task-tracker list done

# Delete a task
./task-tracker delete 2
```
## Project URL
https://roadmap.sh/projects/task-tracker

## Author

Sahil Dhadwal
