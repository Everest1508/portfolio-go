package handlers

import (
	"github.com/everest1508/portfolio/db"
	"github.com/everest1508/portfolio/models"
	"github.com/gin-gonic/gin"
)

func init(){
    db.DBConnect()
}


// GetProjects retrieves all projects
func GetProjects(c *gin.Context) {
    var projects []models.Project
    db.DB.Find(&projects)
    c.JSON(200, projects)
}


func GetProject(c *gin.Context) {
    id := c.Param("id")
    var project models.Project
    if err := db.DB.First(&project, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Record not found"})
        return
    }
    c.JSON(200, project)
}

// CreateProject creates a new project
func CreateProject(c *gin.Context) {
    var project models.Project
    if err := c.ShouldBindJSON(&project); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    if err := db.DB.Create(&project).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to create project"})
        return
    }
    c.JSON(201, project)
}

// UpdateProject updates an existing project
func UpdateProject(c *gin.Context) {
    id := c.Param("id")
    var project models.Project
    if err := db.DB.First(&project, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Record not found"})
        return
    }
    if err := c.ShouldBindJSON(&project); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    if err := db.DB.Save(&project).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to update project"})
        return
    }
    c.JSON(200, project)
}

// DeleteProject deletes a project by ID
func DeleteProject(c *gin.Context) {
    id := c.Param("id")
    var project models.Project
    if err := db.DB.First(&project, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Record not found"})
        return
    }
    if err := db.DB.Delete(&project).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to delete project"})
        return
    }
    c.JSON(204, nil)
}
