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

// @Summary Create Dispetcher
// @Security ApiKeyAuth
// @Tags dispetchers
// @Description creates and adds data about the dispetcher
// @ID create-dispetcher
// @Accept json
// @Produce json
// @Param input body entity.Dispetcher true "dispetcher info"
// @Success 200 {integer} integer "dispetcher ID"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /dispetchers [post]
func (h *Handler) createDispetcher(c *gin.Context) {
	var input entity.Dispetcher
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	dispetcherId, err := h.services.PutlistAPI.CreateDispetcher(context.Background(), input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"dispetcher_id": dispetcherId,
	})
}

type getDispetchersResponse struct {
	Dispetchers []entity.Dispetcher `json:"dispetchers"`
}

// @Summary Get Dispetchers
// @Security ApiKeyAuth
// @Tags dispetchers
// @Description gets data about all dispetchers
// @ID get-dispetchers
// @Accept json
// @Produce json
// @Success 200 {object} getDispetchersResponse "dispetchers info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /dispetchers [get]
func (h *Handler) getDispetchers(c *gin.Context) {
	dispetchers, err := h.services.PutlistAPI.GetDispetchers(context.Background())
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getDispetchersResponse{
		Dispetchers: dispetchers,
	})
}

// @Summary Update Dispetcher
// @Security ApiKeyAuth
// @Tags dispetchers
// @Description updates data about the dispetcher
// @ID update-dispetcher
// @Accept json
// @Produce json
// @Param dispetcher_id  body integer true "dispetcher id"
// @Param updateData body entity.UpdateDispetcherInput true "update dispetcher info"
// @Success 200 {object} statusResponse "updated dispetcher"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /dispetchers/{dispetcher_id} [put]
func (h *Handler) updateDispetcher(c *gin.Context) {
	dispetcherId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var updateData entity.UpdateDispetcherInput

	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	updateMessage, err := h.services.PutlistAPI.UpdateDispetcher(context.Background(), int64(dispetcherId), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "dispetcher not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: updateMessage,
	})
}

// @Summary Delete Dispetcher
// @Security ApiKeyAuth
// @Tags dispetchers
// @Description deletes data about dispetcher
// @ID delete-dispetcher
// @Accept json
// @Produce json
// @Param dispetcher_id body integer true "dispetcher id"
// @Success 200 {object} statusResponse "deleted"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /dispetchers/{dispetcher_id} [delete]
func (h *Handler) deleteDispetcher(c *gin.Context) {
	dispetcherId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	delMessage, err := h.services.PutlistAPI.DeleteDispetcher(context.Background(), int64(dispetcherId))
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "dispetcher not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: delMessage,
	})
}
