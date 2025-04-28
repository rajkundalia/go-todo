package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreateAt  time.Time `json:"created_at"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

// SaveToFile saves tasks to a JSON file
// (tl *Tasklist) is a reciever i.e. method will be associated to
// TaskList
func (tl *TaskList) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromFile loads tasks from a JSON file
func (tl *TaskList) LoadFromFile(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			tl.Tasks = []Task{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, tl)
}

func (tl *TaskList) AddTask(title string) Task {
	newID := 1
	if len(tl.Tasks) > 0 {
		newID = tl.Tasks[len(tl.Tasks)-1].ID + 1
	}

	task := Task{
		ID:        newID,
		Title:     title,
		Completed: false,
		CreateAt:  time.Now(),
	}

	tl.Tasks = append(tl.Tasks, task)
	return task
}

func (tl *TaskList) ListTasks() {
	if len(tl.Tasks) == 0 {
		fmt.Println("No tasks to display!")
		return
	}

	fmt.Println("ID | Completed | Title")
	fmt.Println("----------------------")

	for _, task := range tl.Tasks {
		status := " "
		if task.Completed {
			status = "✓"
		}
		fmt.Printf("%2d | [%s]       | %s\n", task.ID, status, task.Title)
	}
}

func (tl *TaskList) MarkComplete(id int) error {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks[i].Completed = true
			return nil
		}
	}
	return fmt.Errorf("Task with ID %d not found", id)
}

func (tl *TaskList) DeleteTask(id int) error {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = slices.Delete(tl.Tasks, i, i+1)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}
