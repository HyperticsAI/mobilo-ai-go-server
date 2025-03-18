package controllers

import (
	models "go-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Setting struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type CreateSettingRequest struct {
	OrganizationID string    `json:"organization_id" binding:"required"`
	Settings       []Setting `json:"settings" binding:"required"`
}

type UpdateSettingRequest struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func CreateOrganizationSetting(c *gin.Context) {
	var request CreateSettingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, setting := range request.Settings {
		setting := models.OrganizationSetting{
			OrganizationID: request.OrganizationID,
			Key:            setting.Key,
			Value:          setting.Value,
		}
		if err := models.DB.Create(&setting).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusCreated, request.Settings)
}

func GetOrganizationSettings(c *gin.Context) {
	organizationId := c.Param("organizationId")
	if organizationId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Organization ID is required"})
		return
	}

	var settings []models.OrganizationSetting
	if err := models.DB.Where(&models.OrganizationSetting{OrganizationID: organizationId}).Find(&settings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, settings)
}

func GetOrganizationSetting(c *gin.Context) {
	organizationId := c.Param("organizationId")
	if organizationId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Organization ID is required"})
		return
	}

	key := c.Param("key")

	var setting models.OrganizationSetting
	if err := models.DB.First(&setting, "organization_id = ? AND key = ?", organizationId, key).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Setting not found"})
		return
	}

	c.JSON(http.StatusOK, setting)
}

func UpdateOrganizationSetting(c *gin.Context) {
	organizationId := c.Param("organizationId")
	if organizationId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Organization ID is required"})
		return
	}

	var request UpdateSettingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var setting models.OrganizationSetting
	if err := models.DB.First(&setting, "organization_id = ? AND key = ?", organizationId, request.Key).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Setting not found"})
		return
	}

	setting.Value = request.Value

	if err := models.DB.Save(&setting).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, setting)
}
