# Task Tracker

A simple command-line task management tool built in Go. This project allows users to add, list, update, and delete tasks, with data persisted in a JSON file.

## Features

- Add new tasks with descriptions
- List all tasks with their status
- Update task descriptions
- Mark tasks as in-progress, done, or todo
- Delete tasks
- Data persistence using JSON file storage

## Installation

### Prerequisites

- Go 1.16 or later installed on your system. You can download it from [golang.org](https://golang.org/dl/).

### Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/task-tracker.git
   cd task-tracker
   ```

2. Install dependencies (if any):

   ```bash
   go mod tidy
   ```

3. Build the project:

   ```bash
   go build -o task-tracker main.go
   ```

4. Run the application:
   ```bash
   ./task-tracker
   ```

## Usage

The tool uses subcommands to perform operations. Here are the available commands:

- Add a task:

  ```bash
  ./task-tracker add -task "Your task description"
  ```

- List all tasks:

  ```bash
  ./task-tracker list
  ```

- List tasks by status (e.g., done, todo, in-progress):

  ```bash
  ./task-tracker list -list "done"
  ```

- Update a task:

  ```bash
  ./task-tracker update -id 1 -task "Updated task description"
  ```

- Mark a task as done:

  ```bash
  ./task-tracker mark -id 1 -mark "done"
  ```

- Mark a task as in-progress:

  ```bash
  ./task-tracker mark -id 1 -mark "in-progress"
  ```

- Mark a task as todo:

  ```bash
  ./task-tracker mark -id 1 -mark "todo"
  ```

- Delete a task:
  ```bash
  ./task-tracker delete -id 1
  ```

Replace `1` with the actual task ID.

## Topics Learned

This project was created to understand and practice the following topics in Go programming:

- Basic Go syntax and structure
- Structs and JSON marshaling/unmarshaling
- File I/O operations (reading and writing files)
- Command-line argument parsing using the `flag` package
- Time handling with the `time` package
- Error handling and logging
- Working with slices and maps
- JSON data persistence

## Future Plans

This project is a work in progress. Future enhancements may include:

- Adding more task statuses
- Implementing task priorities
- Adding due dates and reminders
- Creating a web interface
- Supporting multiple users or projects
- Improving error handling and validation
- Adding unit tests

Contributions and suggestions are welcome!
