package task

import (
	"fmt"
	"os"
)

type JSONTaskRepo struct {
	jsonFile *os.File
}

func save(*Task) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	fmt.Print(homeDir)
	return nil
}

func (repo *JSONTaskRepo) Add(task *Task) error {
	save(task)
	fmt.Print(repo.jsonFile)
	return nil
}

func (repo *JSONTaskRepo) Update(task *Task) error {

	return nil
}

func (repo *JSONTaskRepo) Delete(id int) error {

	return nil
}

func (repo *JSONTaskRepo) GetAll() ([]*Task, error) {

	return nil, nil
}
