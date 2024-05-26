package main

import (
	"github.com/everest1508/portfolio/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.BioRoutes(r)
	routes.SkillRoutes(r)
	routes.ContactRoutes(r)
	routes.ProjectRoutes(r)
	r.Run(":8080")
}