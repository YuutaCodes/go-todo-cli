package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

// Check if provided index is not out of bounds
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("index out of range")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos
	err := todos.validateIndex(index)
	if err != nil {
		return err
	}

	// split at the index, then attach the items after the index
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	t := &(*todos)[index]

	if t.CompletedAt == nil {
		completionTime := time.Now()
		t.CompletedAt = &completionTime
	} else {
		t.CompletedAt = nil
	}

	return nil
}
