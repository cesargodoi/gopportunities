package handler

import (
	"fmt"
	"net/http"

	"github.com/cesargodoi/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(
			c, http.StatusBadRequest,
			errParamIsRequired("id", "queryParameter").Error(),
		)
		return
	}

	opening := schemas.Opening{}

	// Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(
			c, http.StatusNotFound,
			fmt.Sprintf("opening with id: %v not found", id),
		)
		return
	}

	// Delete Opening
	if err := db.Delete(&opening, id).Error; err != nil {
		sendError(
			c, http.StatusInternalServerError,
			fmt.Sprintf("error deleting opening with id: %v", id),
		)
		return
	}

	sendSuccess(c, "delete-opening", opening)
}
