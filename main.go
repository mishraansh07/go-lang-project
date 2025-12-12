package main

import "fmt"

type Student struct {
	name string
	task string
	day  int
}

func main() {

	// Create two students
	stud1 := Student{
		name: "Ansh Mishra",
		task: "I want to work on my maths project",
		day:  23,
	}

	stud2 := Student{
		name: "Rahul Sharma",
		task: "I want to complete my Go project",
		day:  12,
	}

	// Slice of Students
	studentsList := []Student{stud1, stud2}

	// Print them
	for _, s := range studentsList {
		fmt.Println("Name:", s.name)
		fmt.Println("Task:", s.task)
		fmt.Println("Day:", s.day)
		fmt.Println("------------------")
	}
}
