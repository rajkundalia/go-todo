# Go Todo CLI

A simple command-line Todo application written in Go for Windows.

## Overview

Go Todo CLI is a command-line application that helps you manage your tasks efficiently. This lightweight tool is perfect for developers and power users who prefer terminal-based productivity tools.

## Features

* **Add new tasks:** Easily add new items to your todo list.
* **List all tasks:** View a clear overview of all your pending and completed tasks.
* **Mark tasks as completed:** Update the status of your tasks once they are done.
* **Delete tasks:** Remove tasks that are no longer relevant.
* **Persistent storage with JSON:** Your tasks are saved and loaded using a JSON file, ensuring your data persists between sessions.

## Installation

### Prerequisites

* **Go 1.18 or higher installed on your Windows machine**
    * Download from [golang.org](https://golang.org/dl/)
    * Follow the installation instructions provided on the Go website.

### Setup

1.  **Clone the repository or download the source code:**
    ```bash
    git clone [https://github.com/rajkundalia/go-todo](https://github.com/yourusername/go-todo)
    ```

2.  **Change to the project directory:**
    ```bash
    cd go-todo
    ```

3.  **Build the application:**
    ```bash
    go build -o todo.exe
    ```
    This command will compile the Go source code and create an executable file named `todo.exe` in the project directory.

## Usage

### Running the Application

The application can be run directly from the Command Prompt or PowerShell. Ensure you are in the directory containing the `todo.exe` file, or add the directory to your system's PATH environment variable to run it from anywhere.

### Commands

Here are the available commands:

* **List all tasks:**
    ```bash
    todo.exe list
    ```
    This command displays all the tasks in your todo list, indicating their ID and completion status.

* **Add a new task:**
    ```bash
    todo.exe add "Your new task description here"
    ```
    Replace `"Your new task description here"` with the actual task you want to add. The task description should be enclosed in double quotes if it contains spaces.

* **Mark a task as completed:**
    ```bash
    todo.exe complete <task_id>
    ```
    Replace `<task_id>` with the numerical ID of the task you want to mark as completed. You can find the task ID when listing all tasks. For example:
    ```bash
    todo.exe complete 1
    ```

* **Delete a task:**
    ```bash
    todo.exe delete <task_id>
    ```
    Replace `<task_id>` with the numerical ID of the task you want to delete. You can find the task ID when listing all tasks. For example:
    ```bash
    todo.exe delete 2
    ```