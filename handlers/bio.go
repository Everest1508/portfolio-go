package handlers

import (
	"github.com/everest1508/portfolio/db"
	"github.com/everest1508/portfolio/models"
	"github.com/gin-gonic/gin"
)

// GetBios retrieves all bios
func GetBios(c *gin.Context) {
    var bios []models.Bio
    db.DB.Find(&bios)
    c.JSON(200, bios)
}

// GetBio retrieves a single bio by ID
func GetBio(c *gin.Context) {
    id := c.Param("id")
    var bio models.Bio
    if err := db.DB.First(&bio, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Record not found"})
        return
    }
    c.JSON(200, bio)
}

// CreateBio creates a new bio
func CreateBio(c *gin.Context) {
    var bio models.Bio
    if err := c.ShouldBindJSON(&bio); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    if err := db.DB.Create(&bio).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to create bio"})
        return
    }
    c.JSON(201, bio)
}

// UpdateBio updates an existing bio
func UpdateBio(c *gin.Context) {
    id := c.Param("id")
    var bio models.Bio
    if err := db.DB.First(&bio, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Record not found"})
        return
    }
    if err := c.ShouldBindJSON(&bio); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    if err := db.DB.Save(&bio).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to update bio"})
        return
    }
    c.JSON(200, bio)
}

// DeleteBio deletes a bio by ID
func DeleteBio(c *gin.Context) {
    id := c.Param("id")
    var bio models.Bio
    if err := db.DB.First(&bio, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Record not found"})
        return
    }
    if err := db.DB.Delete(&bio).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to delete bio"})
        return
    }
    c.JSON(204, nil)
}
