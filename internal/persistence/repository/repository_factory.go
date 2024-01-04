package repository

import (
	. "Orgzz-Tasks/internal/models"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Save(newTask *Task) (Task, error) {
	var savedTask Task
	result := r.db.Create(newTask).Scan(&savedTask)

	return savedTask, result.Error
}

func (r *repository) FindAll() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}
