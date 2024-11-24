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

// @Summary Create Putlist
// @Security ApiKeyAuth
// @Tags putlists
// @Description creates and adds data about the putlist
// @ID create-putlist
// @Accept json
// @Produce json
// @Param input body entity.Putlist true "putlist info"
// @Success 200 {integer} integer "putlist ID"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /putlists [post]
func (h *Handler) createPutlist(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input entity.Putlist
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	putlistId, err := h.services.PutlistAPI.CreatePutlist(context.Background(), userId, input)
	if err != nil {
		if errors.Is(err, service.ErrPutlistExists) {
			NewErrorResponse(c, http.StatusBadRequest, service.ErrPutlistExists.Error())
			return
		}
		if errors.Is(err, service.ErrInvalidDateTimeFormat) {
			NewErrorResponse(c, http.StatusBadRequest, service.ErrInvalidDateTimeFormat.Error())
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"putlist_id": putlistId,
	})

}

type getPutlistsResponse struct {
	Putlists []entity.Putlist `json:"putlists"`
}

// @Summary Get Putlists
// @Security ApiKeyAuth
// @Tags putlists
// @Description gets data about all putlist headers
// @ID get-putlists
// @Accept json
// @Produce json
// @Success 200 {object} getPutlistsResponse "putlist headers info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /putlists [get]
func (h *Handler) getPutlists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	putlists, err := h.services.PutlistAPI.GetPutlists(context.Background(), userId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getPutlistsResponse{
		Putlists: putlists,
	})

}

type getPutlistsByNumberResponse struct {
	Putlist entity.Putlist `json:"putlist"`
}

// @Summary Get Putlist By Number
// @Security ApiKeyAuth
// @Tags putlists
// @Description gets data about putlist
// @ID get-putlist-by-number
// @Accept json
// @Produce json
// @Param number path integer true "number"
// @Success 200 {object} getPutlistsByNumberResponse "putlist header info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /putlists/{number} [get]
func (h *Handler) getPutlistByNumber(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid number param")
		return
	}

	putlist, err := h.services.PutlistAPI.GetPutlistByNumber(context.Background(), userId, int64(number))
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "putlist by this number not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getPutlistsByNumberResponse{
		Putlist: putlist,
	})
}

// @Summary Update Putlist
// @Security ApiKeyAuth
// @Tags putlists
// @Description updates data about the putlist header
// @ID update-putlist
// @Accept json
// @Produce json
// @Param number path integer true "number"
// @Param updateData body entity.UpdatePutlistHeaderInput true "update putlist info"
// @Success 200 {object} statusResponse "updated putlist"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /putlists/{number} [put]
func (h *Handler) updatePutlist(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid number param")
		return
	}

	var updateData entity.UpdatePutlistHeaderInput
	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updateMessage, err := h.services.PutlistAPI.UpdatePutlist(context.Background(), userId, int64(number), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "putlist not found")
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

// @Summary Delete Putlist
// @Security ApiKeyAuth
// @Tags putlists
// @Description deletes data about putlist
// @ID delete-putlist
// @Accept json
// @Produce json
// @Param number path integer true "number"
// @Success 200 {object} statusResponse "deleted"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /putlists/{number} [delete]
func (h *Handler) deletePutlist(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid number param")
		return
	}

	delMessage, err := h.services.PutlistAPI.DeletePutlist(context.Background(), userId, int64(number))
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "putlist not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: delMessage,
	})
}
