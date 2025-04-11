package main

import (
	"encoding/json"
	"os"
)

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func saveTodos(todos []Todo) error {
	file, err := os.Create("todos.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(todos)
}

func loadTodos() ([]Todo, error) {
	file, err := os.Open("todos.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var todos []Todo
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
} 