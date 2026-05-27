package task

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type TaskManager struct {
	tasks []Task
}

func NewTaskManager() TaskManager {
	return TaskManager{}
}

func (manager *TaskManager) AddTask(task Task) {
	manager.tasks = append(manager.tasks, task)
}

func (manager *TaskManager) RemoveTask(id string) error {
	for i, task := range manager.tasks {
		if id == task.id {
			manager.tasks = slices.Delete(manager.tasks, i, i+1)
			return nil
		}
	}

	return fmt.Errorf("Error ID:%v not found", id)
}

func (manager *TaskManager) RemoveTasks(tasksStr string) error {
	specialCase := strings.Contains(tasksStr, ":")

	if specialCase {
		split := strings.Split(tasksStr, ":")
		start, errStart := strconv.Atoi(split[0])
		end, errEnd := strconv.Atoi(split[1])

		if errStart != nil || errEnd != nil {
			return fmt.Errorf("Failed to convert string to int (start-end)")
		}

		if start < end {
			for i := start; i <= end; i++ {
				manager.RemoveTask(fmt.Sprint(i))
			}
		} else if start > end {
			for i := start; i >= end; i-- {
				manager.RemoveTask(fmt.Sprint(i))
			}
		} else if start == end {
			manager.RemoveTask(fmt.Sprint(start))
		}
	} else {
		ids := strings.Split(tasksStr, " ")

		for _, id := range ids {
			manager.RemoveTask(id)
		}
	}

	return nil
}

func (manager *TaskManager) containsId(id string) bool {
	for _, task := range manager.tasks {
		if task.id == id {
			return true
		}
	}

	return false
}

func (manager *TaskManager) GetUniqueId() string {
	currentIdInt := 1
	currentId := "1"

	for {
		if manager.containsId(currentId) {
			currentIdInt++
		} else {
			break
		}

		currentId = fmt.Sprint(currentIdInt)
	}

	return currentId
}
