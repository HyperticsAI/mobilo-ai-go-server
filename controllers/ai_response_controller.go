package controllers

import (
	"go-server/helpers"
	models "go-server/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// type CreateAIResponseRequest struct {
// 	OrganizationID string `json:"organization_id" binding:"required"`
// 	Query          string `json:"query" binding:"required"`
// 	Response       string `json:"response" binding:"required"`
// }

// Request represents the incoming API request
type GenerateMessagesRequest = helpers.AiContext

// Response represents the API response
type GenerateMessagesResponse struct {
	Success    bool                     `json:"success"`
	Messages   []helpers.ChannelMessage `json:"messages,omitempty"`
	Error      string                   `json:"error,omitempty"`
	TimeTaken  string                   `json:"time_taken,omitempty"`
	UsedTokens int64                    `json:"used_tokens,omitempty"`
}

type CreateAIResponseFeedbackRequest struct {
	Feedback string `json:"feedback" binding:"required"`
}

func CreateAIResponse(c *gin.Context) {
	// var request CreateAIResponseRequest
	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// response := models.AIResponse{
	// 	OrganizationID: request.OrganizationID,
	// 	Query:          request.Query,
	// 	Response:       request.Response,
	// }
	// if err := models.DB.Create(&response).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// Validate request body
	var input GenerateMessagesRequest
	// TODO: Validate request body & parse it to AiContext
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, GenerateMessagesResponse{
			Success: false,
			Error:   "Invalid request: " + err.Error(),
		})
		return
	}
	// Convert request to BusinessContext
	// input := helpers.AiContext{
	// 	Channel:           req.Channel,
	// 	AdditionalContext: req.AdditionalContext,
	// 	BusinessInfo:      req.BusinessInfo,
	// 	Goal:              req.Goal,
	// 	CustomerProfile:   req.CustomerProfile,
	// }

	// input := helpers.AiContext{
	// 	Channel:           "LinkedIn",
	// 	AdditionalContext: "I want to book a product demo where I will be show them the product and how it works",
	// 	BusinessInfo: helpers.BusinessInfoStruct{
	// 		CompanyName:  "MobiloCard",
	// 		Industry:     "Tech",
	// 		CoreProducts: []string{"MobiloCard Pro", "MobiloCard Business"},
	// 		ValueProps:   []string{"Increase efficiency", "Reduce costs", "Increase productivity", "Increase customer satisfaction"},
	// 	},
	// 	Goal: helpers.GoalStruct{
	// 		Type:        "sales",
	// 		Description: "Book product demo",
	// 		Target:      "Schedule 15-minute call",
	// 	},
	// 	CustomerProfile: helpers.CustomerProfileStruct{
	// 		Name:      "John Doe",
	// 		Title:     "CTO",
	// 		Company:   "Target Corp",
	// 		Industry:  "Retail",
	// 		Interests: []string{"AI", "Digital Transformation", "Customer retention", "Customer satisfaction"},
	// 	},
	// }

	result, err := helpers.GenerateAIResponse(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Time taken to generate messages: %v", result.TimeTaken)
	log.Printf("Tokens used: %v", result.UsedTokens)

	c.JSON(http.StatusCreated, result)
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
