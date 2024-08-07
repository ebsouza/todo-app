package tasks

import (
	"gorm.io/gorm"

	"errors"

	"github.com/google/uuid"

	"time"
)

const (
	taskCreationErrorMessage     string = "Task could not be created"
	taskNotFoundErrorMessage     string = "Task not found"
	taskNotRecoveredErrorMessage string = "Tasks could not be recovered"
	taskNotRemovedErrorMessage   string = "Task could not be removed"
	taskNotUpdatedErrorMessage   string = "Task could not be updated"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	db.AutoMigrate(&Task{})
	return &repository{db}
}

func (r repository) AddTask(task *Task) (uuid.UUID, error) {
	if result := r.DB.Create(task); result.Error != nil {
		return task.ID, errors.New(taskCreationErrorMessage)
	}

	return task.ID, nil
}

func (r repository) GetTask(id string) (*Task, error) {
	task := &Task{}

	if result := r.DB.First(task, "id = ?", id); result.Error != nil {
		return task, errors.New(taskNotFoundErrorMessage)
	}

	return task, nil
}

func (r repository) GetAllTasks(limit int, offset int, status string) ([]Task, error) {
	tasks := []Task{}

	if result := r.DB.Where(&Task{Status: status}).Order("created_at asc").Limit(limit).Offset(offset).Find(&tasks); result.Error != nil {
		return tasks, errors.New(taskNotRecoveredErrorMessage)
	}

	return tasks, nil
}

func (r repository) DeleteTask(id string) (*Task, error) {
	task, err := r.GetTask(id)

	if err != nil {
		return &Task{}, errors.New(taskNotFoundErrorMessage)
	}

	if result := r.DB.Delete(task); result.Error != nil {
		return task, errors.New(taskNotRemovedErrorMessage)
	}

	return task, nil
}

func (r repository) UpdateTask(id string, title string, description string, status string) (*Task, error) {
	task, err := r.GetTask(id)

	if err != nil {
		return &Task{}, errors.New(taskNotFoundErrorMessage)
	}

	task.Title = title
	task.Description = description
	task.Status = status
	task.UpdatedAt = time.Now()

	if result := r.DB.Save(task); result.Error != nil {
		return task, errors.New(taskNotUpdatedErrorMessage)
	}

	return task, nil
}
