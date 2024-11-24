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

// @Summary Create Putlist Body
// @Security ApiKeyAuth
// @Tags  putlist-bodies
// @Description creates and adds data about the putlist body
// @ID create-putlist-body
// @Accept json
// @Produce json
// @Param number path integer true "putlist number"
// @Param input body entity.PutlistBody true "putlist body info"
// @Success 200 {integer} integer "putlist body ID"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{number}/putlist_bodies [post]
func (h *Handler) createPutlistBody(c *gin.Context) {
	putlistNumber, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid putlist_number param")
	}

	var input entity.PutlistBody
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	putlistBodyId, err := h.services.PutlistAPI.CreatePutlistBody(context.Background(), int64(putlistNumber), input)
	if err != nil {
		if errors.Is(err, service.ErrInvalidDateTimeFormat) {
			NewErrorResponse(c, http.StatusBadRequest, service.ErrInvalidDateTimeFormat.Error())
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"putlist_body_id": putlistBodyId,
	})
}

type getPutlistBodiesResponse struct {
	PutlistBodies []entity.PutlistBody `json:"putlist_bodies"`
}

// @Summary Get Putlist Bodies
// @Security ApiKeyAuth
// @Tags  putlist-bodies
// @Description gets data about all putlist bodies
// @ID get-putlist-bodies
// @Accept json
// @Produce json
// @Param number path integer true "putlist number"
// @Success 200 {object} getPutlistBodiesResponse "putlist bodies info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{number}/putlist_bodies [get]
func (h *Handler) getPutlistBodies(c *gin.Context) {
	putlistNumber, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid putlist_number param")
	}
	putlistBodies, err := h.services.PutlistAPI.GetPutlistBodies(context.Background(), int64(putlistNumber))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getPutlistBodiesResponse{
		PutlistBodies: putlistBodies,
	})
}

// @Summary Update Putlist Body
// @Security ApiKeyAuth
// @Tags  putlist-bodies
// @Description updates data about the putlist body
// @ID update-putlist-body
// @Accept json
// @Produce json
// @Param number path integer true "putlist number"
// @Param putlist_body_id path integer true "putlist body id"
// @Param updateData body entity.UpdatePutlistBodyInput true "update putlist body info"
// @Success 200 {object} statusResponse "updated putlist body"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{number}/putlist_bodies/{putlist_body_id} [put]
func (h *Handler) updatePutlistBody(c *gin.Context) {
	putlistBodyId, err := strconv.Atoi(c.Param("putlist_body_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid putlist_body_number param")
	}

	var updateData entity.UpdatePutlistBodyInput
	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updateMessage, err := h.services.PutlistAPI.UpdatePutlistBody(context.Background(), int64(putlistBodyId), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "putlist body not found")
			return
		}
		if errors.Is(err, service.ErrInvalidDateTimeFormat) {
			NewErrorResponse(c, http.StatusBadRequest, service.ErrInvalidDateTimeFormat.Error())
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: updateMessage,
	})
}

// @Summary Delete Putlist Body
// @Security ApiKeyAuth
// @Tags putlist-bodies
// @Description deletes data about putlist body
// @ID delete-putlist-body
// @Accept json
// @Produce json
// @Param number path integer true "putlist number"
// @Param putlist_body_id path integer true "putlist body id"
// @Success 200 {object} statusResponse "deleted"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{number}/putlist_bodies/{putlist_body_id} [delete]
func (h *Handler) deletePutlistBody(c *gin.Context) {
	putlistBodyId, err := strconv.Atoi(c.Param("putlist_body_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid putlist_body_number param")
	}

	delMessage, err := h.services.PutlistAPI.DeletePutlistBody(context.Background(), int64(putlistBodyId))
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "putlist body not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: delMessage,
	})
}
