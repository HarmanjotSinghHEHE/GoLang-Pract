/*Author: Harmanjot Singh 3eee.
Company Name: Coditas
*/

package main

import "fmt"

type Todo struct {
	ID          int
	Title       string
	Description string
	Completion  bool
}

var todos []Todo // Slice to store multiple todos.

func addingTask(title string, description string) error {
	if title == " " {
		return fmt.Errorf("Title Can nOT BE eMPTY")
	}
	id := len(todos) + 1 // ID for the Task
	newTask := Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Completion:  false,
	}
	todos = append(todos, newTask)
	fmt.Println("Task Added")
	return nil
}

func asCompleted(id int) {
	for i, task := range todos {
		if task.ID == id {
			todos[i].Completion = true
			fmt.Println("Task Marked As Completed")
			return
		}
	}
	fmt.Println("Task Not Found")
}

func deletingTask(id int) {
	for i, task := range todos {
		if task.ID == id {
			todos = append(todos[:i], todos[i+1:]...) // Correctly joining slices
			fmt.Printf("Task with ID %d has been deleted.\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", id)
}

func displayTasks() {
	if len(todos) == 0 {
		fmt.Println("No tasks to display.")
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

func main() {
	// Adding tasks
	var ti string
	var dc string
	var tskrem int
	var chstat int

	fmt.Println("*****************************************************************Welcome to ToDo List*****************************************************************")
	fmt.Println("=======> 1.Add a new Task\n==========> 2.Remove task\n==============> 3.Display exisiting tasks\n=================> 4.Mark a task as Completed")
	var key int = 1
	fmt.Println("************************************************************Enter your choice index number*****************************************************************")

	switch key {
	case 1:
		fmt.Println("\nEnter the task name")
		fmt.Scanln(&ti)
		fmt.Println("\nEnter the Description")
		fmt.Scanln(&dc)
		addingTask(ti, dc)
	case 2:
		fmt.Println("\nFor removing a task , enter the task ID")
		fmt.Scanln(&tskrem)
		deletingTask(tskrem)
	case 3:
		displayTasks()
	case 4:
		fmt.Println("\nEnter the task ID , for which you want to mark as completed")
		fmt.Scanln(&chstat)
		asCompleted(chstat)
	default:
		fmt.Println("Enter a valid choice")
		return
	}
	/*
	   Testcases

	   	addingTask(" Learn Go", "Study Go programming basics")
	   	addingTask(" Build a To-Do app", "Create a basic To-Do app in Go")
	   	addingTask(" Write blog post", "Write a blog post about Go")
	   	//Display tasks before deleting
	   	displayTasks()
	   	//Deleting task with ID 2
	   	deletingTask(2)
	   	// Display tasks after deleting
	   	displayTasks()
	*/
}
