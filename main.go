package main

import (
	"flag"
	"fmt"
	"todo/task"
	"todo/utils"
)

func main() {
	addTask := flag.String("add", "", "Add task")
	removeTasks := flag.String("del", "", "Remove task by ID")
	csvPath := flag.String("path", "todo.csv", "Save path file")

	flag.Parse()

	manager := task.NewTaskManager()

	if utils.PathExists(*csvPath) {
		err := manager.ReadFromFile(*csvPath)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	if *addTask != "" {
		newTask := task.NewTaskNow(manager.GetUniqueId(), *addTask)
		manager.AddTask(newTask)

		manager.SaveToFile(*csvPath)
	} else if *removeTasks != "" {
		manager.RemoveTasks(*removeTasks)

		manager.SaveToFile(*csvPath)
	}

	manager.Print()
}
