package task

import "time"

type Task struct {
	id       string
	contents string
	created  int64
}

func NewTask(id string, contents string, created int64) Task {
	return Task{id, contents, created}
}

func NewTaskNow(id string, contents string) Task {
	return Task{id, contents, time.Now().Unix()}
}
