# Go CLI ToDo App
A CLI-based ToDo manager built in Go to practice struct modeling, slice manipulation, and user input handling.

This is my first Go project.

It is a simple command-line ToDo application built to practice Go fundamentals such as:
- Structs
- Slices
- Loops and switch statements
- User input handling
- Error handling and validation

## Features

- View tasks
- Add new tasks
- Mark tasks as completed
- Delete tasks
- Input validation to prevent crashes

## Technical Details

Tasks are stored in-memory using slices.  
User input is converted from 1-based indexing (user view) to 0-based indexing (Go slice indexing).

## How to Run

```bash
go run main.go
