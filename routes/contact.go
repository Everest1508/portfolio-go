package routes

import (
	"github.com/everest1508/portfolio/handlers"
	"github.com/gin-gonic/gin"
)

func ContactRoutes(r *gin.Engine) {
	r.GET("/contact", handlers.GetContacts)
	r.POST("/contact", handlers.CreateContact)
	r.GET("/contact/:id", handlers.GetContact)
	r.DELETE("/contact/:id", handlers.DeleteContact)
	r.PUT("/contact/:id", handlers.UpdateContact)
}