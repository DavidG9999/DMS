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

// @Summary Create Contragent
// @Security ApiKeyAuth
// @Tags contragents
// @Description creates data about the contragent
// @ID create-contragent
// @Accept json
// @Produce json
// @Param input body entity.Contragent true "contragent info"
// @Success 200 {integer} integer "contragent ID"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /contragents [post]
func (h *Handler) createContragent(c *gin.Context) {
	var input entity.Contragent
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	contragentId, err := h.services.PutlistAPI.CreateContragent(context.Background(), input)
	if err != nil {
		if errors.Is(err, service.ErrContragentExists) {
			NewErrorResponse(c, http.StatusBadRequest, service.ErrContragentExists.Error())
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"contragent_id": contragentId,
	})
}

type getContragentsResponse struct {
	Contragents []entity.Contragent `json:"contragents"`
}

// @Summary Get Contragents
// @Security ApiKeyAuth
// @Tags contragents
// @Description gets data about all contragents
// @ID get-contragents
// @Accept json
// @Produce json
// @Success 200 {object} getContragentsResponse "contragents info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /contragents [get]
func (h *Handler) getContragents(c *gin.Context) {
	contragents, err := h.services.PutlistAPI.GetContragents(context.Background())
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getContragentsResponse{
		Contragents: contragents,
	})
}

// @Summary Update Contragent
// @Security ApiKeyAuth
// @Tags contragents
// @Description updates data about the contragent
// @ID update-contragent
// @Accept json
// @Produce json
// @Param contragent_id body integer true "contragent id"
// @Param updateData body entity.UpdateContragentInput true "update contragent info"
// @Success 200 {object} statusResponse "updated contragent"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /contragents/{contragent_id} [put]
func (h *Handler) updateContragent(c *gin.Context) {
	contragentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var updateData entity.UpdateContragentInput

	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	updateMessage, err := h.services.PutlistAPI.UpdateContragent(context.Background(), int64(contragentId), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "contragent not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: updateMessage,
	})
}

// @Summary Delete Contragent
// @Security ApiKeyAuth
// @Tags contragents
// @Description deletes data about contragent
// @ID delete-contragent
// @Accept json
// @Produce json
// @Param contragent_id path integer true "contragent id"
// @Success 200 {object} statusResponse "deleted"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /contragents/{contragent_id} [delete]
func (h *Handler) deleteContragent(c *gin.Context) {
	contragentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	delMessage, err := h.services.PutlistAPI.DeleteContragent(context.Background(), int64(contragentId))
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "contragent not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: delMessage,
	})
}
