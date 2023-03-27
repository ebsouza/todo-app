package repository

import (
	"errors"
	"todo/server/data"
	"todo/server/domain"
)

func AddTask(task domain.Task) error {
	data.Tasks = append(data.Tasks, task)

	return nil
}

func GetTask(task_id string) (domain.Task, error) {
	all_tasks, _ := GetAllTasks()

	for _, task := range all_tasks {
		if task.ID == task_id {
			return task, nil
		}
	}

	return domain.Task{}, errors.New("Task not found")
}

func GetAllTasks() ([]domain.Task, error) {
	return data.Tasks, nil
}

func RemoveTask(task_id string) (domain.Task, error) {
	var found bool
	var index int
	var task domain.Task

	all_tasks, _ := GetAllTasks()

	for i, t := range all_tasks {
		if t.ID == task_id {
			found = true
			index = i
			task = t
			break
		}
	}

	if !found {
		return task, errors.New("Task not found")
	}

	data.Tasks[index] = data.Tasks[len(data.Tasks)-1]
	data.Tasks[len(data.Tasks)-1] = domain.Task{}
	data.Tasks = data.Tasks[:len(data.Tasks)-1]

	return task, nil
}
