package todo

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Todo struct {
	Id       uint32
	Task     string
	Date     string
	Deadline string
	Done     bool
}

func Add(task string, deadline string) Todo {
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	id := rand.Uint32()

	var todo Todo
	todo.Id = id
	todo.Task = task
	todo.Deadline = deadline
	todo.Done = false
	todo.Date = date

	file, err := os.OpenFile("./data.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	singleRow := []string{
		strconv.FormatUint(uint64(todo.Id), 10),
		todo.Task,
		todo.Date,
		todo.Deadline,
		strconv.FormatBool(todo.Done),
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(singleRow); err != nil {
		log.Fatalf("Error writing single row: %s", err)
	}

	return todo
}

func GetAll() []Todo {
	file, err := os.Open("./data.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV: %s", err)
	}

	todos := []Todo{}

	for _, row := range records {
		var todo Todo

		u64, err := strconv.ParseUint(row[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		todoId := uint32(u64)

		todoDone, err := strconv.ParseBool(row[4])
		if err != nil {
			log.Fatal(err)
		}

		todo.Id = todoId
		todo.Task = row[1]
		todo.Date = row[2]
		todo.Deadline = row[3]
		todo.Done = todoDone

		todos = append(todos, todo)
	}

	return todos
}

func GetTodos() []Todo {
	file, err := os.Open("./data.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV: %s", err)
	}

	todos := []Todo{}

	for _, row := range records {
		var todo Todo

		u64, err := strconv.ParseUint(row[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		todoId := uint32(u64)

		todoDone, err := strconv.ParseBool(row[4])
		if err != nil {
			log.Fatal(err)
		}

		todo.Id = todoId
		todo.Task = row[1]
		todo.Date = row[2]
		todo.Deadline = row[3]
		todo.Done = todoDone

		if !todo.Done {
			todos = append(todos, todo)
		}
	}

	return todos
}

func CompleteTodo(id uint32) {
	todos := GetAll()

	for i, todo := range todos {
		if todo.Id == id {
			todos[i].Done = true
		}
	}

	newFile, err := os.Create("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	writer := csv.NewWriter(newFile)
	if err := writer.WriteAll(getTodosInString(todos)); err != nil {
		log.Fatalf("Error writing single row: %s", err)
	}
}

func getTodosInString(todos []Todo) [][]string {
	var todosStr [][]string
	for _, todo := range todos {
		singleRow := []string{
			strconv.FormatUint(uint64(todo.Id), 10),
			todo.Task,
			todo.Date,
			todo.Deadline,
			strconv.FormatBool(todo.Done),
		}
		todosStr = append(todosStr, singleRow)
	}
	return todosStr
}
