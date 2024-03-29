package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
)

func NewList() (*List, error) {
	l := &List{TaskList: make(map[int]Task)}
	err := l.GetTasks()
	if err != nil {
		return l, err
	}
	return l, nil
}

func (l *List) GetTasks() error {
	data, err := os.ReadFile("list.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &l.TaskList)
}

func (l *List) AddNew(task Task) {
	l.TaskList[len(l.TaskList)+1] = task
	l.Save()
}

func (l *List) Complete(id int) error {
	value, ok := l.TaskList[id]
	if !ok {
		return errors.New("can't find a task with this id")
	}
	value.Completed = "\u2713"
	l.TaskList[id] = value
	l.Save()
	fmt.Printf("Well done for completing task n.%d !\n", id)
	return nil
}

func (l *List) Save() error {
	json, err := json.MarshalIndent(l.TaskList, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile("list.json", []byte(json), 0644)
}

func (l *List) Read() {
	var sortedList []int
	for key := range l.TaskList {
		sortedList = append(sortedList, key)
	}
	sort.Ints(sortedList)
	for _, id := range sortedList {
		fmt.Printf("%d. %s [%s] \n", id, l.TaskList[id].Task, l.TaskList[id].Completed)
	}
}

func (l *List) Delete(id int) error {
	if _, ok := l.TaskList[id]; !ok {
		return errors.New("can't find a task with this id")
	}
	delete(l.TaskList, id)
	for index, value := range l.TaskList {
		if index > id {
			l.TaskList[index-1] = value
			delete(l.TaskList, index)
		}
	}
	l.Save()
	return nil
}
