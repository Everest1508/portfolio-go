package routes

import (
	"github.com/everest1508/portfolio/handlers"
	"github.com/gin-gonic/gin"
)

func SkillRoutes(r *gin.Engine) {
	r.GET("/skill", handlers.GetSkills)
	r.POST("/skill", handlers.CreateSkill)
	r.GET("/skill/:id", handlers.GetSkill)
	r.DELETE("/skill/:id", handlers.DeleteSkill)
	r.PUT("/skill/:id", handlers.UpdateSkill)
}