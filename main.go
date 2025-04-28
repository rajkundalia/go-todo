package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

const taskFile = "tasks.json"

func main() {
	var taskList TaskList
	err := taskList.LoadFromFile(taskFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCmd.String("title", "", "Title of the task")

	comleteCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	completeID := comleteCmd.Int("id", 0, "ID of the task to mark complete")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteID := deleteCmd.Int("id", 0, "ID of the task to delete")

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		taskList.ListTasks()

	case "add":
		err := addCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing add command %v\n", err)
			os.Exit(1)
		}

		if *addTitle == "" {
			if addCmd.NArg() > 0 {
				*addTitle = addCmd.Arg(0)
			} else {
				fmt.Println("Error: Title is required")
				addCmd.Usage()
				os.Exit(1)
			}
		}

		task := taskList.AddTask(*addTitle)
		fmt.Printf("Added task %d: %s\n", task.ID, task.Title)
		taskList.SaveToFile(taskFile)

	case "complete":
		err := comleteCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing complete command: %v\n", err)
			os.Exit(1)
		}

		id := *completeID
		if id == 0 && comleteCmd.NArg() > 0 {
			idArg, err := strconv.Atoi(comleteCmd.Arg(0))
			if err != nil {
				fmt.Println("Error: Invalid task ID")
				os.Exit(1)
			}
			id = idArg
		}

		if id == 0 {
			fmt.Println("Error: Task ID is required")
			comleteCmd.Usage()
			os.Exit(1)
		}

		err = taskList.MarkComplete(id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Marked task %d as completed\n", id)
		taskList.SaveToFile(taskFile)

	case "delete":
		err := deleteCmd.Parse(os.Args[2:])

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing delete command: %v\n", err)
			os.Exit(1)
		}
		id := *deleteID
		if id == 0 && deleteCmd.NArg() > 0 {
			idArg, err := strconv.Atoi(deleteCmd.Arg(0))
			if err != nil {
				fmt.Println("Error: Invalid task ID")
				os.Exit(1)
			}
			id = idArg
		}

		if id == 0 {
			fmt.Println("Error: Task ID is required")
			deleteCmd.Usage()
			os.Exit(1)
		}

		err = taskList.DeleteTask(id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Deleted task %d\n", id)
		taskList.SaveToFile(taskFile)

	default:
		printUsage()
		os.Exit(1)
	}

}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  todo list                       - List all tasks")
	fmt.Println("  todo add -title \"Task title\"    - Add a new task")
	fmt.Println("  todo add \"Task title\"           - Add a new task (alternative)")
	fmt.Println("  todo complete -id <id>          - Mark a task as completed")
	fmt.Println("  todo complete <id>              - Mark a task as completed (alternative)")
	fmt.Println("  todo delete -id <id>            - Delete a task")
	fmt.Println("  todo delete <id>                - Delete a task (alternative)")
}
