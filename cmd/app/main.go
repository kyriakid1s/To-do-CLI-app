package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Task struct {
	Task      string `json:"task"`
	Completed string `json:"completed"`
}

type List struct {
	TaskList map[int]Task `json:"task_list"`
}

func main() {
	var task string
	var read bool
	var completed, deleted int
	var help bool
	flag.StringVar(&task, "task", "", "Add a task")
	flag.IntVar(&completed, "complete", 0, "Mark a task as completed")
	flag.BoolVar(&read, "read", false, "Get the list")
	flag.IntVar(&deleted, "delete", 0, "Delete the selected Id")
	flag.BoolVar(&help, "help", false, "Show Commands")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("A parameter is required. Type -help to show the commands.")
		os.Exit(0)
	}
	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	list, err := NewList()
	if err != nil {
		log.Fatal(err)
	}
	switch {
	case task != "":
		list.AddNew(Task{Task: task, Completed: " "})
		fmt.Println("Task added!")
	case read:
		list.Read()
	case completed != 0:
		err := list.Complete(completed)
		if err != nil {
			log.Print(err)
		}
	case deleted != 0:
		err := list.Delete(deleted)
		if err != nil {
			log.Println(err)
		}
	}
}
