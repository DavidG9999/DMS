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

// @Summary Create Organization
// @Security ApiKeyAuth
// @Tags organizations
// @Description creates and adds data about the organization
// @ID create-organization
// @Accept json
// @Produce json
// @Param input body entity.Organization true "organization info"
// @Success 200 {integer} integer "organization ID"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /organizations [post]
func (h *Handler) createOrganization(c *gin.Context) {
	var input entity.Organization
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	organizationId, err := h.services.PutlistAPI.CreateOrganization(context.Background(), input)
	if err != nil {
		if errors.Is(err, service.ErrOrganizationExists) {
			NewErrorResponse(c, http.StatusBadRequest, service.ErrOrganizationExists.Error())
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"organization_id": organizationId,
	})
}

type getOrganizationsResponse struct {
	Organizations []entity.Organization `json:"organizations"`
}

// @Summary Get Organizations
// @Security ApiKeyAuth
// @Tags organizations
// @Description gets data about all organizations
// @ID get-organizations
// @Accept json
// @Produce json
// @Success 200 {object} getOrganizationsResponse "organizations info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /organizations [get]
func (h *Handler) getOrganizations(c *gin.Context) {
	organizations, err := h.services.PutlistAPI.GetOrganizations(context.Background())
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getOrganizationsResponse{
		Organizations: organizations,
	})
}

// @Summary Update Organizaion
// @Security ApiKeyAuth
// @Tags organizations
// @Description updates data about the organization
// @ID update-organization
// @Accept json
// @Produce json
// @Param organization_id body integer true "organization id"
// @Param updateData body entity.UpdateOrganizationInput true "update organization info"
// @Success 200 {object} statusResponse "updated organization"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /organizations/{organization_id} [put]
func (h *Handler) updateOrganization(c *gin.Context) {
	organiationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var updateData entity.UpdateOrganizationInput

	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	updateMessage, err := h.services.PutlistAPI.UpdateOrganization(context.Background(), int64(organiationId), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "organization not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: updateMessage,
	})
}

// @Summary Delete Organization
// @Security ApiKeyAuth
// @Tags organizations
// @Description deletes data about organization
// @ID delete-organization
// @Accept json
// @Produce json
// @Param organization_id body integer true "organization id"
// @Success 200 {object} statusResponse "deleted"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /organizations/{organization_id} [delete]
func (h *Handler) deleteOrganization(c *gin.Context) {
	organiationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	delMessage, err := h.services.PutlistAPI.DeleteOrganization(context.Background(), int64(organiationId))
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "organization not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: delMessage,
	})
}
