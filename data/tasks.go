package data

import (
	"errors"

	"github.com/google/uuid"
)

const (
	TodoStatus = "todo"
	DoneStatus = "done"
)

type Task struct {
	Title       string
	Description string
	Status      string
}

type TasksDb struct {
	TasksMap map[string]*Task
	IdList   []string
}

func (tdb *TasksDb) DeleteTask(id string) error {
	_, ok := tdb.TasksMap[id]
	if !ok {
		return errors.New("task with id " + id + " not found")
	}
	idx := -1
	for i, idInList := range tdb.IdList {
		if idInList == id {
			idx = i
		}
	}
	if idx == -1 {
		return errors.New("task map and id list out of sync")
	}
	delete(tdb.TasksMap, id)
	tdb.IdList = append(tdb.IdList[:idx], tdb.IdList[idx+1:]...)
	return nil
}

func (tdb *TasksDb) AddTask(id string, task *Task) {
	tdb.TasksMap[id] = task
	tdb.IdList = append(tdb.IdList, id)
}

func (t *Task) ToggleStatus() {
	if t.Status == TodoStatus {
		t.Status = DoneStatus
		return
	}
	t.Status = TodoStatus
}

func NewTask(title, description string) *Task {
	t := new(Task)
	t.Title = title
	t.Description = description
	t.Status = TodoStatus
	return t
}

var Tasks TasksDb

func Load() {
	Tasks = TasksDb{
		TasksMap: make(map[string]*Task),
		IdList:   make([]string, 0),
	}
	sampleTasks := []*Task{
		{Title: "Get haircut", Description: "Need to get a haircut its about time", Status: TodoStatus},
		{Title: "Car wash", Description: "Need to get the car washed", Status: TodoStatus},
		{Title: "Apply for visa", Description: "Bunch of stuff happening here", Status: TodoStatus},
		{Title: "Learn htmx", Description: "Learning htmx is going to save my career", Status: TodoStatus},
		{Title: "Interview prep", Description: "Prepping using leetcode and that", Status: TodoStatus},
		{Title: "Grocery shopping", Description: "Lets do some shopping man", Status: DoneStatus},
		{Title: "Get new work desk", Description: "Need a new desk cause old one is too tall", Status: DoneStatus},
	}
	for _, task := range sampleTasks {
		id := uuid.NewString()
		Tasks.AddTask(id, task)
	}
}
