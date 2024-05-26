package routes

import (
	"github.com/everest1508/portfolio/handlers"
	"github.com/gin-gonic/gin"
)

func BioRoutes(r *gin.Engine){
	r.GET("/bio",handlers.GetBios)
	r.POST("/bio",handlers.CreateBio)
	r.GET("/bio/:id",handlers.GetBio)
	r.DELETE("/bio/:id",handlers.DeleteBio)
	r.PUT("/bio/:id",handlers.UpdateBio)
}