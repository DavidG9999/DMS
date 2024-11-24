package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/service"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}

// @Summary Get User
// @Security ApiKeyAuth
// @Tags user
// @Description gives information about the user
// @ID get-user-info
// @Accept json
// @Produce json
// @Success 200 {object} UserResponse "user"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user [get]
func (h *Handler) getUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	userName, userEmail, err := h.services.User.GetUserById(context.Background(), userId)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "user not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, UserResponse{
		UserName: userName,
		Email:    userEmail,
	})
}

// @Summary Update Name
// @Security ApiKeyAuth
// @Tags user
// @Description updates username
// @ID update-name
// @Accept json
// @Produce json
// @Param updateData body entity.UpdateNameUserInput true "update password"
// @Success 200 {object} statusResponse "updated name"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/name [put]
func (h *Handler) updateName(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var updateData entity.UpdateNameUserInput
	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updateMessage, err := h.services.User.UpdateName(context.Background(), userId, updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "user not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: updateMessage,
	})
}

// @Summary Update Password
// @Security ApiKeyAuth
// @Tags user
// @Description updates user password
// @ID update-password
// @Accept json
// @Produce json
// @Param updateData body entity.UpdatePasswordUserInput true "update password"
// @Success 200 {object} statusResponse "updated password"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/password [put]
func (h *Handler) updatePassword(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var updateData entity.UpdatePasswordUserInput
	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updateMessage, err := h.services.User.UpdatePassword(context.Background(), userId, updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "user not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: updateMessage,
	})

}

// @Summary Delete User
// @Security ApiKeyAuth
// @Tags user
// @Description deletes user account
// @ID delete-user-account
// @Accept json
// @Produce json
// @Success 200 {object} statusResponse "deleted"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	delMessage, err := h.services.User.DeleteUser(context.Background(), userId)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "user not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: delMessage,
	})

}
