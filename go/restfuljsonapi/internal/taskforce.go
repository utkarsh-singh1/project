package taskforce

import (
	"fmt"
	"time"
)

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

// TaskStore is a in-memory database that will store tasks
type TaskStore struct {
	tasks  map[int]Task
	nextID int
}

// New will create a new instance of taskstore to store each task with a new Id starting with 0
func New() *TaskStore {

	ts := &TaskStore{}

	ts.tasks = make(map[int]Task)
	ts.nextID = 0

	return ts

}

// CreateTask creates a new task upon api calling
func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {

	task := Task{
		Id:   ts.nextID,
		Text: text,
		Due:  due,
	}

	task.Tags = make([]string, len(tags))

	copy(task.Tags, tags)

	ts.tasks[ts.nextID] = task

	ts.nextID++

	return task.Id

}

// GetTask gets the task based on the id provided
func (ts *TaskStore) GetTask(id int) (Task, error) {

	t, ok := ts.tasks[id]
	if ok {

		return t, nil

	} else {

		return Task{}, fmt.Errorf("No task assign to particular ID = %d ", id)
	}

}

// DeleteTask delete task by ID
func (ts *TaskStore) DeleteTask(id int) error {

	if _, ok := ts.tasks[id]; !ok {
		return fmt.Errorf("No task assign to particular ID = %d", id)
	}

	delete(ts.tasks, id)

	return nil

}

// DeleteAllTask delete all task in the db
func (ts *TaskStore) DeleteAllTask() error {

	ts.tasks = make(map[int]Task)

	return nil
}

// GetAllTask get all task present in the db
func (ts *TaskStore) GetAllTask() []Task {

	allTask := make([]Task, 0, len(ts.tasks))

	for _, task := range ts.tasks {

		allTask = append(allTask, task)

	}

	return allTask

}

// GetTaskByTags get the task by tags
func (ts *TaskStore) GetTaskByTag(tag string) []Task {

	var tasks []Task

	for _, task := range ts.tasks {

		for _, taskTag := range task.Tags {

			if taskTag == tag {
				tasks = append(tasks, task)
			}
		}

	}

	return tasks

}

// GetTaskByDueDate get the task by due date
func (ts *TaskStore) GetTaskByDueDate(year int, month time.Month, date int) []Task {

	var tasks []Task

	for _, task := range ts.tasks {

		y, m, d := task.Due.Date()
		if y == year && m == month && d == date {
			tasks = append(tasks, task)
		}
	}

	return tasks

}
