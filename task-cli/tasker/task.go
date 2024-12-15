package tasker

import (
	"encoding/json"
	"fmt"
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

func Add(task *Task) (uint, error) {
	tasks, err := readTasks()
	if err != nil {
		return 0, err
	}

	var highest uint

	if len(tasks) > 0 {
		highest = tasks[len(tasks)-1].Id
	}

	task.Id = highest + 1

	tasks = append(tasks, task)

	err = writeTasks(tasks)
	if err != nil {
		return 0, err
	}

	return task.Id, nil
}


func Update(id uint, value string, update string) error {

	tasks, err := readTasks()
	if err != nil {
		return err
	}

	switch update {
	case "status":
		for _, v := range tasks {
			if v.Id == id {
				v.Status = value
				v.UpdatedAt = time.Now()
				break
			}
		}
	case "description":
		for _, v := range tasks {
			if v.Id == id {
				v.Description = value
				v.UpdatedAt = time.Now()
				break
			}
		}
	default:
		return fmt.Errorf("invalid update value: %s", update)
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
