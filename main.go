package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Title string
	Done  bool
}

// func showMenu(reader bufio.Reader) string {
// 	fmt.Print(" Welcome to Go ToDo App \n\n")

// 	fmt.Print(" 1 - View Tasks : \n")
// 	fmt.Print(" 2 - Add Tasks : \n")
// 	fmt.Print(" 3 - Mark Task Done : \n")
// 	fmt.Print(" 4 - Delete Task : \n")
// 	fmt.Print(" 5 - Exit : \n\n")

// 	fmt.Print(" Enter option number : ")
// 	optionReader, _ := reader.ReadString('\n')
// 	optionReader = strings.TrimSpace(optionReader)
// 	return optionReader
// }

func loadWork(fileName string) ([]Task, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	return tasks, err
}

func showMenu(reader *bufio.Reader) string {
	for {
		fmt.Print("\n\n")
		fmt.Println("----------------------")
		fmt.Println("Welcome to Go ToDo App")
		fmt.Print("----------------------\n")

		fmt.Println("1 - View Tasks")
		fmt.Println("2 - Add Task")
		fmt.Println("3 - Mark Task Done")
		fmt.Println("4 - Delete Task")
		fmt.Println("5 - Exit")
		fmt.Println()

		fmt.Print("Enter option number: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Try again.")
			continue
		}

		option := strings.TrimSpace(input)

		switch option {
		case "1", "2", "3", "4", "5":
			return option
		default:
			fmt.Print("Invalid option. Please choose 1–5.\n\n")
		}
	}
}

// func viewTasks(tasks []Task) {
// 	for i, task := range tasks {
// 		fmt.Printf("%d - %s - %t\n", i, task.Title, task.Done)
// 	}
// }

func viewTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Print("No tasks available.\n\n")
		return
	}

	fmt.Print("\n\n")
	fmt.Println("\nYour Tasks:")
	fmt.Print("---------------------------\n\n")

	for i, task := range tasks {
		status := "❌ "
		if task.Done {
			status = "✅ "
		}

		fmt.Printf("%d. %s [%s]\n", i+1, task.Title, status)
	}

	fmt.Println()
}

func addTask(tasks []Task, reader *bufio.Reader) []Task {
	fmt.Print("\n\n")
	fmt.Println("\nAdd Your Tasks:")
	fmt.Print("---------------------------\n\n")

	fmt.Print("Enter task title: ")
	titleInput, _ := reader.ReadString('\n')
	title := strings.TrimSpace(titleInput)

	if title == "" {
		fmt.Println("Task title cannot be empty")
		return tasks
	}

	fmt.Print("Is task done? (true/false): ")
	doneInput, _ := reader.ReadString('\n')
	doneInput = strings.TrimSpace(doneInput)

	isDone, err := strconv.ParseBool(doneInput)
	if err != nil {
		fmt.Println("Invalid input. Defaulting to 'false'")
		isDone = false
	}

	newTask := Task{
		Title: title,
		Done:  isDone,
	}

	tasks = append(tasks, newTask)
	fmt.Print("Task added successfully!\n\n")

	return tasks
}

func markTaskDone(tasks []Task, reader *bufio.Reader) {
	fmt.Print("3 - Mark Task Done:\n")
	fmt.Print("Enter the task number: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	taskNum, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid Task Number")
		return
	}

	// Subtract 1 to match slice index
	taskNum--

	if taskNum < 0 || taskNum >= len(tasks) {
		fmt.Println("Task does not exist")
		return
	}

	tasks[taskNum].Done = true
	fmt.Println("Task marked done!")
}

func saveWork(fileName string, tasks []Task) error {
	fmt.Print("Saving the Tasks List...\n")
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)

}

func deleteTask(tasks []Task, reader *bufio.Reader) []Task {
	fmt.Print("\n\n")
	fmt.Print("4 - Delete Task:\n")
	fmt.Print("Enter the task number: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	taskNum, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid task number")
		return tasks
	}

	// If displaying tasks starting from 1, uncomment:
	taskNum--

	if taskNum < 0 || taskNum >= len(tasks) {
		fmt.Println("Task does not exist")
		return tasks
	}

	tasks = append(tasks[:taskNum], tasks[taskNum+1:]...)
	fmt.Println("Task deleted successfully!")

	return tasks
}

func main() {
	tasks := []Task{
		{Title: "Learn Go", Done: false},
		{Title: "Build App", Done: true},
	}

	const fileName = "tasks.json"

	tasks, err := loadWork(fileName)
	if err != nil {
		fmt.Println("No existing tasks found. Starting fresh.")
		tasks = []Task{}
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		option := showMenu(reader)

		switch option {
		case "1":
			viewTasks(tasks)

		case "2":
			tasks = addTask(tasks, reader)
			saveWork(fileName, tasks)

		case "3":
			markTaskDone(tasks, reader)
			saveWork(fileName, tasks)

		case "4":
			tasks = deleteTask(tasks, reader)
			saveWork(fileName, tasks)

		case "5":
			fmt.Println("Goodbye 👋")
			return
		}

		fmt.Println("----------------------")
	}
}

// func main() {

// taskName := "Learn Go"
// isDone := false

// tasks := []string{"Task 1", "Task 2", "Task 3"}

// type Task struct {
// 	Title string
// 	Done  bool
// }
//tasks := []Task{Task{Title: "Learn Go", Done: false}, Task{Title: "Build App", Done: true}}

//fmt.Println("Welcome to Go ToDo App") // prints with new line
// fmt.Print("Welcome to Go ToDo App")   // prints without new line
// fmt.Printf("Welcome to Go ToDo App")  // formatted  Uses format specifiers like: %s → string   %d → integer    %f → float %t -> boolean

// fmt.Println()
// fmt.Printf("Task :  %s   |   Done : %t \n", taskName, isDone)

// for i, task := range tasks {
// 	fmt.Printf("%d - %s - %t\n", i, task.Title, task.Done)
// }

// for _, task := range tasks {
// 	fmt.Println(task)
// }

// fmt.Println("Add Task : ")

// reader := bufio.NewReader(os.Stdin)

// fmt.Println("Add task title : ")

// taskTitle, _ := reader.ReadString('\n') // delim → delimiter → tells Go when to stop reading. The function keeps reading characters until it sees that delimiter.

// fmt.Println("Does task done ? (true / false) : ")
// taskDone, _ := reader.ReadString('\n')
// taskTitle = strings.TrimSpace(taskTitle)
// isTaskDone := false

// if taskDone == "true" {
// 	isTaskDone = true
// }
// newTask := Task{Title: taskTitle, Done: isTaskDone}
// tasks = append(tasks, newTask)
// for i, task := range tasks {
// 	fmt.Printf("%d - %s - %t\n", i, task.Title, task.Done)
// }

// 	for {
// 		reader := bufio.NewReader(os.Stdin)

// 		optionReader := showMenu(reader)

// 		switch optionReader {
// 		case "1":
// 			viewTasks(tasks)
// 		case "2":
// 			tasks = addTask(tasks, reader)

// 		case "3":
// 			markTaskDone(tasks, reader)

// 		case "4":
// 			tasks = deleteTask(tasks, reader)

// 		case "5":
// 			os.Exit(0)

// 		}
// 	}

// }
