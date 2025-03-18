package controllers

import (
	models "go-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateAIResponseRequest struct {
	OrganizationID string `json:"organization_id" binding:"required"`
	Query          string `json:"query" binding:"required"`
	Response       string `json:"response" binding:"required"`
}

type CreateAIResponseFeedbackRequest struct {
	Feedback string `json:"feedback" binding:"required"`
}

func CreateAIResponse(c *gin.Context) {
	var request CreateAIResponseRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := models.AIResponse{
		OrganizationID: request.OrganizationID,
		Query:          request.Query,
		Response:       request.Response,
	}
	if err := models.DB.Create(&response).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response)
}

func GetOrganizationAIResponses(c *gin.Context) {
	organizationId, err := uuid.Parse(c.Param("organizationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	var responses []models.AIResponse
	if err := models.DB.Where(&models.AIResponse{OrganizationID: organizationId.String()}).Find(&responses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, responses)
}

func GetOrganizationAIResponse(c *gin.Context) {
	organizationId, err := uuid.Parse(c.Param("organizationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	var response models.AIResponse
	if err := models.DB.Where(&models.AIResponse{OrganizationID: organizationId.String(), ID: id}).First(&response).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AI response not found"})
		return
	}
	c.JSON(http.StatusOK, response)
}

func CreateAIResponseFeedback(c *gin.Context) {
	organizationId, err := uuid.Parse(c.Param("organizationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	var request CreateAIResponseFeedbackRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	feedback := models.AIResponseFeedback{
		OrganizationID: organizationId.String(),
		AIResponseID:   id,
		Feedback:       request.Feedback,
	}
	if err := models.DB.Create(&feedback).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, feedback)
}
