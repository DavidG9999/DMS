package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/service"

	"github.com/gin-gonic/gin"
)

// @Summary Create Mehanic
// @Security ApiKeyAuth
// @Tags mehanics
// @Description creates and adds data about the mehanic
// @ID create-mehanic
// @Accept json
// @Produce json
// @Param input body entity.Mehanic true "mehanic info"
// @Success 200 {integer} integer "mehanic ID"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /mehanics [post]
func (h *Handler) createMehanic(c *gin.Context) {
	var input entity.Mehanic
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	mehanicId, err := h.services.PutlistAPI.CreateMehanic(context.Background(), input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"mehanic_id": mehanicId,
	})
}

type getMehanicsResponse struct {
	Mehanics []entity.Mehanic `json:"mehanics"`
}

// @Summary Get Mehanics
// @Security ApiKeyAuth
// @Tags mehanics
// @Description gets data about all mehanics
// @ID get-mehanics
// @Accept json
// @Produce json
// @Success 200 {object} getMehanicsResponse "mehanics info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /mehanics [get]
func (h *Handler) getMehanics(c *gin.Context) {
	mehanics, err := h.services.PutlistAPI.GetMehanics(context.Background())
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getMehanicsResponse{
		Mehanics: mehanics,
	})
}

// @Summary Update Mehanic
// @Security ApiKeyAuth
// @Tags mehanics
// @Description updates data about the mehanic
// @ID update-mehanic
// @Accept json
// @Produce json
// @Param mehanic_id body integer true "mehanic id"
// @Param updateData body entity.UpdateMehanicInput true "update mehanic info"
// @Success 200 {object} statusResponse "updated mehanic"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /mehanics/{mehanic_id} [put]
func (h *Handler) updateMehanic(c *gin.Context) {
	mehanicId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var updateData entity.UpdateMehanicInput

	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	updateMessage, err := h.services.PutlistAPI.UpdateMehanic(context.Background(), int64(mehanicId), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound,  "mehanic not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: updateMessage,
	})
}

// @Summary Delete Mehanic
// @Security ApiKeyAuth
// @Tags mehanics
// @Description deletes data about mehanic
// @ID delete-mehanic
// @Accept json
// @Produce json
// @Param mehanic_id body integer true "mehanic id"
// @Success 200 {object} statusResponse "deleted"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /mehanics/{mehanic_id} [delete]
func (h *Handler) deleteMehanic(c *gin.Context) {
	mehanicId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	delMessage, err := h.services.PutlistAPI.DeleteMehanic(context.Background(), int64(mehanicId))
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "mehanic not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: delMessage})
}
