package handlers

import (
	"github.com/everest1508/portfolio/db"
	"github.com/everest1508/portfolio/models"
	"github.com/gin-gonic/gin"
)

// GetContacts retrieves all contacts
func GetContacts(c *gin.Context) {
    var contacts []models.Contact
    db.DB.Find(&contacts)
    c.JSON(200, contacts)
}

// GetContact retrieves a single contact by ID
func GetContact(c *gin.Context) {
    id := c.Param("id")
    var contact models.Contact
    if err := db.DB.First(&contact, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Record not found"})
        return
    }
    c.JSON(200, contact)
}

// CreateContact creates a new contact
func CreateContact(c *gin.Context) {
    var contact models.Contact
    if err := c.ShouldBindJSON(&contact); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    if err := db.DB.Create(&contact).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to create contact"})
        return
    }
    c.JSON(201, contact)
}

// UpdateContact updates an existing contact
func UpdateContact(c *gin.Context) {
    id := c.Param("id")
    var contact models.Contact
    if err := db.DB.First(&contact, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Record not found"})
        return
    }
    if err := c.ShouldBindJSON(&contact); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    if err := db.DB.Save(&contact).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to update contact"})
        return
    }
    c.JSON(200, contact)
}

// DeleteContact deletes a contact by ID
func DeleteContact(c *gin.Context) {
    id := c.Param("id")
    var contact models.Contact
    if err := db.DB.First(&contact, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Record not found"})
        return
    }
    if err := db.DB.Delete(&contact).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to delete contact"})
        return
    }
    c.JSON(204, nil)
}
