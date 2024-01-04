package controllers

import "github.com/gin-gonic/gin"

type HealthCheckController struct {
}

func (controller *HealthCheckController) Test(context *gin.Context) {
	context.JSON(200, gin.H{"status": "OK"})
}
