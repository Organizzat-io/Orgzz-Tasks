package web

import (
	. "Orgzz-Tasks/internal/web/controllers"
	"github.com/gin-gonic/gin"
)

type RouteLoader struct {
}

func (loader *RouteLoader) LoadRoutes() {
	healthCheckController := HealthCheckController{}
	taskController := TaskController{}

	router := gin.Default()
	router.GET("healthCheck", healthCheckController.Test)

	router.GET("tasks", taskController.FindAllTasks)
	router.POST("tasks", taskController.Add)

	err := router.Run()
	if err != nil {
		return
	}
}
