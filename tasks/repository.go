package tasks

import (
	"gorm.io/gorm"

	"errors"

	"github.com/google/uuid"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	db.AutoMigrate(&Task{})
	return &repository{db}
}

func (r repository) AddTask(task *Task) error {
	task.ID = uuid.New()

	if result := r.DB.Create(task); result.Error != nil {
		return errors.New("Task could not be created")
	}

	return nil
}

func (r repository) GetTask(task_id string) (Task, error) {
	var task Task

	if result := r.DB.First(&task, "id = ?", task_id); result.Error != nil {
		return Task{}, errors.New("Task not found")
	}

	return task, nil
}

func (r repository) GetAllTasks() ([]Task, error) {
	var tasks []Task

	if result := r.DB.Find(&tasks); result.Error != nil {
		return []Task{}, errors.New("Tasks could not be recovered")
	}

	return tasks, nil
}

func (r repository) DeleteTask(task_id string) (Task, error) {
	task, _ := r.GetTask(task_id)

	if result := r.DB.Delete(&task); result.Error != nil {
		return task, errors.New("Tasks could not be removed")
	}

	return task, nil
}

func (r repository) UpdateTask(task_id string, task_info Task) (Task, error) {
	task, _ := r.GetTask(task_id)

	task.Title = task_info.Title
	task.Description = task_info.Description
	task.Status = task_info.Status

	if result := r.DB.Save(&task); result.Error != nil {
		return task, errors.New("Tasks could not be updated")
	}

	return task, nil
}
