package main

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
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
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: go run main.go [add/list/done]")
	// 	return
	// }

	// tasks := loadTasks()

	for {
		action := selectAction()
		tasks := loadTasks()

		switch action {
		case "Add Task":
			addTask(tasks)
		case "List Tasks":
			listTasks(tasks)
		case "Mark Task Done":
			markDone(tasks)
		case "Delete Task":
			deleteTask(tasks)
		case "Exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command. Use add/list/done.")
		}
	}
}

func selectAction() string {
	prompt := promptui.Select{
		Label: "Choose an action",
		Items: []string{"Add Task", "List Tasks", "Mark Task Done", "Delete Task", "Exit"},
	}

	_, result, _ := prompt.Run()
	return result
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

func addTask(tasks []Task) {
	// if len(args) == 0 {
	// 	fmt.Println("Please provide a task.")
	// 	return
	// }

	prompt := promptui.Prompt{
		Label: "Enter Task",
	}
	description, _ := prompt.Run()

	if description == "" {
		fmt.Println("Description cannot be empty.")
		return
	}
	// task := args[0]
	task := Task{Description: description, Done: false}
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

func markDone(tasks []Task) {
	// if len(args) == 0 {
	// 	fmt.Println("Please provide a task number.")
	// 	return
	// }
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	// index, err := strconv.Atoi(args[0])

	prompt := promptui.Select{
		Label: "Select task to mark as done",
		Items: getTaskDescriptions(tasks),
	}

	// if err != nil || index < 0 || index >= len(tasks) {
	// 	fmt.Println("Invalid task number.")
	// 	return
	// }
	index, _, _ := prompt.Run()

	tasks[index].Done = true
	saveTasks(tasks)
	fmt.Println("Marked task as done:", tasks[index].Description)
}

func deleteTask(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to delete.")
		return
	}

	// index, err := strconv.Atoi(args[0])
	// if err != nil || index < 0 || index >= len(tasks) {
	// 	fmt.Println("Invalid task number.")
	// 	return
	// }

	prompt := promptui.Select{
		Label: "Select task to delete",
		Items: getTaskDescriptions(tasks),
	}
	index, _, _ := prompt.Run()

	task := tasks[index]
	tasks = append(tasks[:index], tasks[index+1:]...)
	saveTasks(tasks)

	fmt.Println("Deleted task:", task.Description)
}

func getTaskDescriptions(tasks []Task) []string {
	var descriptions []string

	for _, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}
		descriptions = append(descriptions, fmt.Sprintf("%s %s", status, task.Description))
	}

	return descriptions
}
