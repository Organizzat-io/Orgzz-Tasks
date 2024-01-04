package repository

import . "Orgzz-Tasks/internal/models"

type Repository interface {
	Save(newTask *Task) (Task, error)
	FindAll() ([]Task, error)
}
