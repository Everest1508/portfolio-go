package routes

import (
	"github.com/everest1508/portfolio/handlers"
	"github.com/gin-gonic/gin"
)

func ProjectRoutes(r *gin.Engine) {
	r.GET("/project", handlers.GetProjects)
	r.POST("/project", handlers.CreateProject)
	r.GET("/project/:id", handlers.GetProject)
	r.DELETE("/project/:id", handlers.DeleteProject)
	r.PUT("/project/:id", handlers.UpdateProject)
}