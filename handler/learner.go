package handler

import (
	"net/http"
	"ruang_belajar/helper"
	"ruang_belajar/models/learners"
	"ruang_belajar/service"

	"github.com/gin-gonic/gin"
)

type learnerHandler struct {
	learnerService service.LearnerService
}

func NewLearnerHandler(learnerService service.LearnerService) *learnerHandler {
	return &learnerHandler{learnerService}
}

func (h *learnerHandler) FetchLearner(c *gin.Context) {
	currentLearner := c.MustGet("currentLearner").(learners.Learner)
	learnerID := currentLearner.ID

	learner, err := h.learnerService.GetLearnerByID(learnerID)
	if err != nil || learner.ID < 1 {
		message := "Failed to get learner"
		if learner.ID < 1 {
			message = "learner not found"
		}

		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Success fetch learner", http.StatusOK, "success", learners.FormatInfoLearner(learner))
	c.JSON(http.StatusOK, response)
}

func (h *learnerHandler) UpdateLearner(c *gin.Context) {
	var inputID learners.GetLearnerInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update learner", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	learner, err := h.learnerService.GetLearnerByID(inputID.ID)
	if err != nil || learner.ID < 1 {
		message := "Failed to get learner"
		if learner.ID < 1 {
			message = "learner not found"
		}

		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	var inputData learners.CreateLearnerInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update learner", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateLearner, err := h.learnerService.UpdateLearner(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update learner", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update learner", http.StatusOK, "success", learners.FormatInfoLearner(updateLearner))
	c.JSON(http.StatusOK, response)

}
