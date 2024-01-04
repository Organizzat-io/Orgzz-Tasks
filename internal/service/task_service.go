package service

import (
	"Orgzz-Tasks/config"
	. "Orgzz-Tasks/internal/models"
	"Orgzz-Tasks/internal/persistence/repository"
)

func SaveTask(newTask *Task) (Task, error) {
	db := config.OpenDbConnection()
	repo := repository.NewRepository(db)

	result, err := repo.Save(newTask)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	return result, err
}

func FindAllTasks() ([]Task, error) {
	db := config.OpenDbConnection()
	repo := repository.NewRepository(db)

	result, err := repo.FindAll()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	return result, err
}
