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

// @Summary Create Driver
// @Security ApiKeyAuth
// @Tags drivers
// @Description creates and adds data about the driver
// @ID create-driver
// @Accept json
// @Produce json
// @Param input body entity.Driver true "driver info"
// @Success 200 {integer} integer "driver ID"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /drivers [post]
func (h *Handler) createDriver(c *gin.Context) {
	var input entity.Driver
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	driverId, err := h.services.PutlistAPI.CreateDriver(context.Background(), input)
	if err != nil {
		if errors.Is(err, service.ErrDriverExists) {
			NewErrorResponse(c, http.StatusBadRequest, service.ErrDriverExists.Error())
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"driver_id": driverId,
	})
}

type getDriversResponse struct {
	Drivers []entity.Driver `json:"drivers"`
}

// @Summary Get Drivers
// @Security ApiKeyAuth
// @Tags drivers
// @Description gets data about all drivers
// @ID get-drivers
// @Accept json
// @Produce json
// @Success 200 {object} getDriversResponse "drivers info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /drivers [get]
func (h *Handler) getDrivers(c *gin.Context) {
	drivers, err := h.services.PutlistAPI.GetDrivers(context.Background())
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getDriversResponse{
		Drivers: drivers,
	})
}

// @Summary Update Driver
// @Security ApiKeyAuth
// @Tags drivers
// @Description updates data about the driver
// @ID update-driver
// @Accept json
// @Produce json
// @Param driver_id body integer true "driver id"
// @Param updateData body entity.UpdateDriverInput true "update driver info"
// @Success 200 {object} statusResponse "updated driver"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /drivers/{driver_id} [put]
func (h *Handler) updateDriver(c *gin.Context) {
	driverId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var updateData entity.UpdateDriverInput

	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	updateMessage, err := h.services.PutlistAPI.UpdateDriver(context.Background(), int64(driverId), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "driver not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: updateMessage,
	})
}

// @Summary Delete Driver
// @Security ApiKeyAuth
// @Tags drivers
// @Description deletes data about driver
// @ID delete-driver
// @Accept json
// @Produce json
// @Param driver_id body integer true "driver id"
// @Success 200 {object} statusResponse "deleted"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /drivers/{driver_id} [delete]
func (h *Handler) deleteDriver(c *gin.Context) {
	driverId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	delMessage, err := h.services.PutlistAPI.DeleteDriver(context.Background(), int64(driverId))
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "driver not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: delMessage,
	})
}
