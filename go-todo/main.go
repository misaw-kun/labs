package main

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// const filename = "tasks.txt"
const filename = "tasks.json"

// var tasks []string
// var completed = make(map[int]bool)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [add/list/done]")
		return
	}

	tasks := loadTasks()

	switch os.Args[1] {
	case "add":
		addTask(os.Args[2:], tasks)
	case "list":
		listTasks(tasks)
	case "done":
		markDone(os.Args[2:], tasks)
	case "delete":
		deleteTask(os.Args[2:], tasks)
	default:
		fmt.Println("Unknown command. Use add/list/done.")
	}
}

func loadTasks() []Task {
	// file, err := os.Open(filename)
	file, err := os.ReadFile(filename)
	if err != nil {
		return []Task{}
	}
	// defer file.Close()

	var tasks []Task

	// .txt file based tasks
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	parts := strings.SplitN(line, ",", 2)
	// 	if len(parts) != 2 {
	// 		continue
	// 	}
	// 	done, _ := strconv.ParseBool(parts[0])
	// 	tasks = append(tasks, Task{Description: parts[1], Done: done})
	// }

	if err := json.Unmarshal(file, &tasks); err != nil {
		fmt.Println("Error reading tasks:", err)
		return []Task{}
	}
	return tasks
}

func saveTasks(tasks []Task) {
	// file, err := os.Create(filename)
	data, err := json.MarshalIndent(tasks, "", "  ") // pretty printing json
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	// defer file.Close()
	// for _, task := range tasks {
	// 	fmt.Fprintf(file, "%t,%s\n", task.Done, task.Description)
	// }

	os.WriteFile(filename, data, 0644)
}

func addTask(args []string, tasks []Task) {
	if len(args) == 0 {
		fmt.Println("Please provide a task.")
		return
	}
	// task := args[0]
	task := Task{Description: strings.Join(args, " "), Done: false}
	tasks = append(tasks, task)
	saveTasks(tasks)
	fmt.Println("Added task: ", task.Description)
}

func listTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	fmt.Println("Your tasks:")
	for i, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i, status, task.Description)
	}
}

func markDone(args []string, tasks []Task) {
	if len(args) == 0 {
		fmt.Println("Please provide a task number.")
		return
	}
	index, err := strconv.Atoi(args[0])

	if err != nil || index < 0 || index >= len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks[index].Done = true
	saveTasks(tasks)
	fmt.Println("Marked task as done:", tasks[index].Description)
}

func deleteTask(args []string, tasks []Task) {
	if len(args) == 0 {
		fmt.Println("Please provide a task number to delete.")
		return
	}

	index, err := strconv.Atoi(args[0])
	if err != nil || index < 0 || index >= len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	task := tasks[index]
	tasks = append(tasks[:index], tasks[index+1:]...)
	saveTasks(tasks)

	fmt.Println("Deleted task:", task.Description)
}
