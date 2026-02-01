package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"
)

func calcLastIndex() int64 {
	file, err := os.Open("tasks.json")
	if err != nil {
		// file doesn't exist yet → no tasks
		if os.IsNotExist(err) {
			return 0
		}
		log.Fatal(err)
	}
	defer file.Close()

	var lastID int64 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var t Task
		if err := json.Unmarshal(scanner.Bytes(), &t); err != nil {
			log.Fatal(err)
		}
		if t.ID > lastID {
			lastID = t.ID
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lastID
}

func calcIndex() int64 {
	return calcLastIndex() + 1
}

type Task struct {
	ID        int64     `json:"id"`
	Desc      string    `json:"desc"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func WriteTask(desc string) {
	file, err := os.OpenFile(
		"tasks.json",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0o644,
	)
	if err != nil {
		log.Fatal("Err opening/creating file", err)
	}
	defer file.Close()

	data := Task{
		ID:        calcIndex(),
		Desc:      desc,
		Status:    "in-progress",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// marshal struct to json bytes
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error marshalling data")
	}

	// bufio for better performance
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// write json + new line
	_, err = writer.Write(append(jsonData, '\n'))
	if err != nil {
		log.Fatal("error writing to file:", err)
	}
}

func UpdateTask(id int64, newDesc string) {
	file, err := os.Open("tasks.json")
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	var tasks []Task

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var t Task
		if err := json.Unmarshal(scanner.Bytes(), &t); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, t)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Desc = newDesc
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		log.Fatal("task not found")
	}

	// rewrite file
	file, err = os.OpenFile(
		"tasks.json",
		os.O_TRUNC|os.O_WRONLY,
		0o644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, t := range tasks {
		data, err := json.Marshal(t)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := writer.Write(append(data, '\n')); err != nil {
			log.Fatal(err)
		}
	}
}

func DeleteTask(id int64) {
	file, err := os.Open("tasks.json")
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	var tasks []Task

	scanner := bufio.NewScanner(file)

	found := false
	for scanner.Scan() {
		var t Task
		if err := json.Unmarshal(scanner.Bytes(), &t); err != nil {
			log.Fatal(err)
		}
		if t.ID == id {
			found = true
			continue
		}
		tasks = append(tasks, t)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if !found {
		log.Fatal("task not found")
	}

	// rewrite file
	file, err = os.OpenFile(
		"tasks.json",
		os.O_TRUNC|os.O_WRONLY|os.O_CREATE,
		0o644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, t := range tasks {
		data, _ := json.Marshal(t)
		writer.Write(append(data, '\n'))
	}
}

func MarkTask(id int64, mark string) {
	file, err := os.Open("tasks.json")
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	var tasks []Task

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var t Task
		if err := json.Unmarshal(scanner.Bytes(), &t); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, t)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = mark
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		log.Fatal("task not found")
	}

	// rewrite file
	file, err = os.OpenFile(
		"tasks.json",
		os.O_TRUNC|os.O_WRONLY,
		0o644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, t := range tasks {
		data, err := json.Marshal(t)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := writer.Write(append(data, '\n')); err != nil {
			log.Fatal(err)
		}
	}
}

func printTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// header
	fmt.Fprintln(w, "ID\tDESCRIPTION\tSTATUS\tCREATED\tMODIFIED")

	for _, t := range tasks {
		fmt.Fprintf(
			w,
			"%d\t%s\t%s\t%s\t%s\n",
			t.ID,
			t.Desc,
			t.Status,
			t.CreatedAt.Format("02/01 15:04"),
			t.UpdatedAt.Format("02/01 15:04"),
		)
	}
	w.Flush()
}

func ListTask(list string) {
	file, err := os.Open("tasks.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var tasks []Task

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var t Task
		if err := json.Unmarshal(scanner.Bytes(), &t); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, t)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var filtered []Task

	switch list {
	case "", "all":
		filtered = tasks
	case "todo":
		for _, t := range tasks {
			if t.Status == "todo" {
				filtered = append(filtered, t)
			}
		}
	case "doing":
		for _, t := range tasks {
			if t.Status == "in-progress" {
				filtered = append(filtered, t)
			}
		}
	case "done":
		for _, t := range tasks {
			if t.Status == "done" {
				filtered = append(filtered, t)
			}
		}
	default:
		log.Fatal("unknown list filter :%s", list)
	}

	printTasks(filtered)
}

func main() {
	addFlag := flag.NewFlagSet("add", flag.ExitOnError)
	updateFlag := flag.NewFlagSet("update", flag.ExitOnError)
	deleteFlag := flag.NewFlagSet("delete", flag.ExitOnError)
	markFlag := flag.NewFlagSet("mark", flag.ExitOnError)
	listFlag := flag.NewFlagSet("list", flag.ExitOnError)
	_ = listFlag
	_ = addFlag

	// addTask := addFlag.String("task", "", "Add a new task")

	lastIndex := calcLastIndex()
	updateTaskID := updateFlag.Int64("id", lastIndex, "Id of which you want to update task for")
	updateTaskMsg := updateFlag.String("task", "", "updated task")

	deleteTaskID := deleteFlag.Int64("id", lastIndex, "Id of task you want to delete")

	markTaskID := markFlag.Int64("id", lastIndex, "id of which you want to change the task progress for")
	markTask := markFlag.String("mark", "in-progress", "is it in-progress or done")

	// listTask := listFlag.String("list", "", "empty, done, todo, in-progress")

	// check if args specified or not
	if len(os.Args) < 2 {
		fmt.Println("This prog requires additional cmds, please do --help")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		// addFlag.Parse(os.Args[2:])
		// WriteTask(*addTask)
		if len(os.Args) < 3 {
			fmt.Println(`write your "task you want to do"`)
		}
		WriteTask(os.Args[2])

	case "update":
		updateFlag.Parse(os.Args[2:])
		UpdateTask(*updateTaskID, *updateTaskMsg)

	case "delete":
		deleteFlag.Parse(os.Args[2:])
		DeleteTask(*deleteTaskID)

	case "mark":
		markFlag.Parse(os.Args[2:])
		MarkTask(*markTaskID, *markTask)

	case "list":
		var listType string
		if len(os.Args) < 3 {
			listType = ""
		} else {
			listType = os.Args[2]
		}
		ListTask(listType)

	default:
		fmt.Println("no/wrong subcmd entered")
		os.Exit(1)
	}
}
