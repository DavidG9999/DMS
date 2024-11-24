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

// @Summary Create Auto
// @Security ApiKeyAuth
// @Tags autos
// @Description creates and adds new data about the car
// @ID create-auto
// @Accept json
// @Produce json
// @Param input body entity.Auto true "auto info"
// @Success 200 {integer} integer "auto_id"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /autos [post]
func (h *Handler) createAuto(c *gin.Context) {
	var input entity.Auto
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	autoId, err := h.services.PutlistAPI.CreateAuto(context.Background(), input)
	if err != nil {
		if errors.Is(err, service.ErrAutoExists) {
			NewErrorResponse(c, http.StatusBadRequest,  service.ErrAutoExists.Error())
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"auto_id": autoId,
	})

}

type getAutosResponse struct {
	Autos []entity.Auto `json:"autos,omitempty"`
}

// @Summary Get Autos
// @Security ApiKeyAuth
// @Tags autos
// @Description get all data about the cars
// @ID get-autos
// @Accept json
// @Produce json
// @Success 200 {object} getAutosResponse "autos info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /autos [get]
func (h *Handler) getAutos(c *gin.Context) {
	autos, err := h.services.PutlistAPI.GetAutos(context.Background())
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAutosResponse{
		Autos: autos,
	})
}

// @Summary Update Auto
// @Security ApiKeyAuth
// @Tags autos
// @Description updates data about the car
// @ID update-auto
// @Accept json
// @Produce json
// @Param auto_id path integer true "auto id"
// @Param input body entity.UpdateAutoInput true "update auto info"
// @Success 200 {object} statusResponse "updated auto"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /autos/{auto_id} [put]
func (h *Handler) updateAuto(c *gin.Context) {

	autoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input entity.UpdateAutoInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updateMessage, err := h.services.PutlistAPI.UpdateAuto(context.Background(), int64(autoId), input)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "auto not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: updateMessage,
	})

}

// @Summary Delete Auto
// @Security ApiKeyAuth
// @Tags autos
// @Description deletes data about the car
// @ID delete-auto
// @Accept json
// @Produce json
// @Param auto_id path integer true "auto id"
// @Success 200 {object} statusResponse "deleted"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /autos/{auto_id} [delete]
func (h *Handler) deleteAuto(c *gin.Context) {
	autoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	delMessage, err := h.services.PutlistAPI.DeleteAuto(context.Background(), int64(autoId))
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "auto not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: delMessage,
	})

}
