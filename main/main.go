package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"nouman.com/todo"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		switch args[0] {
		case "list":
			showTodos(len(args) > 1 && args[1] == "all")
		case "add":
			var task, deadline string
			fmt.Printf("What's the task?: ")
			scanner := bufio.NewScanner(os.Stdin)

			if scanner.Scan() {
				task = scanner.Text()
			}

			fmt.Printf("What's the deadline? [YYYY-MM-DD]: ")
			_, err := fmt.Scanln(&deadline)
			if err != nil {
				log.Fatal(err)
			}

			todo.Add(task, deadline)
		case "done":
			todos := showTodos(false)
			var completed int
			fmt.Printf("Which of the todos have you done? : ")
			_, err := fmt.Scanln(&completed)
			if err != nil {
				log.Fatal(err)
			}
			todo.CompleteTodo(todos[completed-1].Id)
		}
	}

	// Add a todo
}

func showTodos(all bool) []todo.Todo {
	var todos []todo.Todo
	if all {
		todos = todo.GetAll()
	} else {
		todos = todo.GetTodos()
	}
	for i, todo := range todos {
		var doneStr string
		if todo.Done {
			doneStr = "DONE"
		} else {
			doneStr = "TODO"
		}
		fmt.Printf("[%s] %d. %s\nDeadline: %s\n\n", doneStr, i+1, todo.Task, todo.Deadline)
	}

	return todos
}

