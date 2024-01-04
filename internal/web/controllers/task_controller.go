package controllers

import (
	. "Orgzz-Tasks/internal/models"
	"Orgzz-Tasks/internal/service"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
}

func (controller *TaskController) Add(context *gin.Context) {
	var newTask Task
	if err := context.BindJSON(&newTask); err != nil {
		return
	}

	result, err := service.SaveTask(&newTask)
	if err != nil {
		context.JSON(400, err.Error())
	}
	context.JSON(200, result)
}

func (controller *TaskController) FindAllTasks(context *gin.Context) {
	tasks, err := service.FindAllTasks()
	if err != nil {
		context.JSON(400, err.Error())
	}
	context.JSON(200, tasks)
}
