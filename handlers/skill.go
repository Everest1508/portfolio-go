package handlers

import (
	"github.com/everest1508/portfolio/db"
	"github.com/everest1508/portfolio/models"
	"github.com/gin-gonic/gin"
)

func GetSkills(c *gin.Context) {
    var skills []models.Skill
    db.DB.Find(&skills)
    c.JSON(200, skills)
}

// GetSkill retrieves a single skill by ID
func GetSkill(c *gin.Context) {
    id := c.Param("id")
    var skill models.Skill
    if err := db.DB.First(&skill, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Record not found"})
        return
    }
    c.JSON(200, skill)
}

// CreateSkill creates a new skill
func CreateSkill(c *gin.Context) {
    var skill models.Skill
    if err := c.ShouldBindJSON(&skill); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    if err := db.DB.Create(&skill).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to create skill"})
        return
    }
    c.JSON(201, skill)
}

// UpdateSkill updates an existing skill
func UpdateSkill(c *gin.Context) {
    id := c.Param("id")
    var skill models.Skill
    if err := db.DB.First(&skill, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Record not found"})
        return
    }
    if err := c.ShouldBindJSON(&skill); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    if err := db.DB.Save(&skill).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to update skill"})
        return
    }
    c.JSON(200, skill)
}

// DeleteSkill deletes a skill by ID
func DeleteSkill(c *gin.Context) {
    id := c.Param("id")
    var skill models.Skill
    if err := db.DB.First(&skill, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Record not found"})
        return
    }
    if err := db.DB.Delete(&skill).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to delete skill"})
        return
    }
    c.JSON(204, nil)
}
