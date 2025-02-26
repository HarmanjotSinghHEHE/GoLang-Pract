package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

// Todo represents a task in the ToDo list.
type Todo struct {
	ID          int
	Title       string
	Description string
	Completion  bool
}

var (
	todos  []Todo     // Slice to store multiple todos.
	mu     sync.Mutex // Mutex to handle concurrent access to the todos slice.
	nextID = 1        // Track the next ID for new tasks.
)

// Error definitions
var (
	ErrEmptyTitleOrDescription = errors.New("title or description cannot be empty")
	ErrTaskNotFound            = errors.New("task not found")
)

// addingTask adds a new task to the todos slice.
func addingTask(title string, description string) error {
	mu.Lock()
	defer mu.Unlock()

	// Validate inputs
	if title == "" || description == "" {
		return ErrEmptyTitleOrDescription
	}

	// Create a new task
	newTask := Todo{
		ID:          nextID,
		Title:       title,
		Description: description,
		Completion:  false,
	}
	todos = append(todos, newTask)
	nextID++

	log.Println("Task Added: ", newTask)
	return nil
}

// asCompleted marks a task as completed by its ID.
func asCompleted(id int) error {
	mu.Lock()
	defer mu.Unlock()

	// Search for the task by ID
	for i, task := range todos {
		if task.ID == id {
			todos[i].Completion = true
			log.Println("Task Marked As Completed: ", task)
			return nil
		}
	}

	return ErrTaskNotFound
}

// deletingTask deletes a task by its ID.
func deletingTask(id int) error {
	mu.Lock()
	defer mu.Unlock()

	// Search for the task by ID
	for i, task := range todos {
		if task.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			log.Printf("Task with ID %d has been deleted.\n", id)
			return nil
		}
	}

	return ErrTaskNotFound
}

// displayTasks displays all tasks in the todos slice.
func displayTasks() {
	mu.Lock()
	defer mu.Unlock()

	if len(todos) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("All tasks:")
	for _, task := range todos {
		completionStatus := "Not Completed"
		if task.Completion {
			completionStatus = "Completed"
		}
		fmt.Printf("ID: %d | Title: %s | Description: %s | Status: %s\n", task.ID, task.Title, task.Description, completionStatus)
	}
}

// runCLI handles the command-line interface for the ToDo app.
func runCLI() {
	for {
		fmt.Println("\n*****************************************************************Welcome to ToDo List*****************************************************************")
		fmt.Println("=======> 1. Add a new Task\n==========> 2. Remove task\n==============> 3. Display existing tasks\n=================> 4. Mark a task as Completed\n====================> 5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			log.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			var title, description string
			fmt.Print("Enter the task name: ")
			fmt.Scanln(&title)
			fmt.Print("Enter the description: ")
			fmt.Scanln(&description)

			err := addingTask(title, description)
			if err != nil {
				log.Println("Error adding task:", err)
			}

		case 2:
			var id int
			fmt.Print("Enter the task ID to delete: ")
			fmt.Scanln(&id)

			err := deletingTask(id)
			if err != nil {
				log.Println("Error deleting task:", err)
			}

		case 3:
			displayTasks()

		case 4:
			var id int
			fmt.Print("Enter the task ID to mark as completed: ")
			fmt.Scanln(&id)

			err := asCompleted(id)
			if err != nil {
				log.Println("Error marking task as completed:", err)
			}

		case 5:
			fmt.Println("Exiting the ToDo app. Goodbye!")
			return

		default:
			log.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}

func main() {
	// Start the CLI
	runCLI()
}
