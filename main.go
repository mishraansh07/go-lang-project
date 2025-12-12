package main

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (t Todos) validateIndex(index int) error {
	if index < 0 || index >= len(t) {
		return errors.New("index out of range")
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	t := *todos
	todo := &t[index]

	if !todo.Completed {
		completedTime := time.Now()
		todo.CompletedAt = &completedTime
	} else {
		todo.CompletedAt = nil
	}

	todo.Completed = !todo.Completed
	(*todos)[index] = *todo

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(
			strconv.Itoa(index),
			t.Title,
			completed,
			t.CreatedAt.Format(time.RFC1123),
			completedAt,
		)
	}

	table.Render()
}

func main() {
	todos := Todos{}
	todos.add("Buy groceries")
	todos.add("Study Go Lang")
	todos.toggle(1)
	todos.print()
}
