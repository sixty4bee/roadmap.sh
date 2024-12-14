package tasker

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"time"
)

type Task struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func readTasks() ([]*Task, error) {
	tasks := []*Task{}

	tasksData, err := os.ReadFile(filepath.Join("tasks.json"))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(tasksData, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func writeTasks(tasks []*Task) error {
	tasksData, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join("tasks.json"), tasksData, 0777)
	if err != nil {
		return err
	}

	return nil
}

func Add(task *Task) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	var highest uint
	for _, v := range tasks {
		if v.Id > highest {
			highest = v.Id
		}
	}

	task.Id = highest + 1

	tasks = append(tasks, task)

	return writeTasks(tasks)
}

func Update(id uint, description string) error {

	tasks, err := readTasks()
	if err != nil {
		return err
	}

	for _, v := range tasks {
		if v.Id == id {
			v.Description = description
			break
		}
	}

	return writeTasks(tasks)
}

func Delete(id uint) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	tasks = slices.DeleteFunc(tasks, func(t *Task) bool {
		return t.Id == id
	})

	return writeTasks(tasks)
}

func GetAll() ([]*Task, error) {
	return readTasks()
}
