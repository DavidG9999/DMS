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

// @Summary Create Bank Account
// @Security ApiKeyAuth
// @Tags bank-accounts
// @Description creates new data about the bank account
// @ID create-bank-account
// @Accept json
// @Produce json
// @Param organization_id path int true "organization id"
// @Param input body entity.BankAccount true "bank account info"
// @Success 200 {integer} integer "bank account ID"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{organization_id}/bank_accounts [post]
func (h *Handler) createBankAccount(c *gin.Context) {
	organiationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid organization_id param")
	}

	var input entity.BankAccount
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	bankAccountId, err := h.services.PutlistAPI.CreateBankAccount(context.Background(), int64(organiationId), input)
	if err != nil {
		if errors.Is(err, service.ErrBankAccExists) {
			NewErrorResponse(c, http.StatusBadRequest, service.ErrBankAccExists.Error())
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"bank_account_id": bankAccountId,
	})
}

type getBankAccountsResponse struct {
	BankAccounts []entity.BankAccount `json:"bank_accounts"`
}

// @Summary Get Bank Accounts
// @Security ApiKeyAuth
// @Tags bank-accounts
// @Description gets data about all bank accounts
// @ID get-bank-accounts
// @Accept json
// @Produce json
// @Param organization_id path int true "organization id"
// @Success 200 {object} getBankAccountsResponse "get bank accounts info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{organization_id}/bank_accounts [get]
func (h *Handler) getBankAccounts(c *gin.Context) {
	organizationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid organization_id param")
	}
	bankAccounts, err := h.services.PutlistAPI.GetBankAccounts(context.Background(), int64(organizationId))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getBankAccountsResponse{
		BankAccounts: bankAccounts,
	})
}

// @Summary Update Bank Account
// @Security ApiKeyAuth
// @Tags bank-accounts
// @Description updates data about the bank account
// @ID update-bank-account
// @Accept json
// @Produce json
// @Param organization_id path int true "organization id"
// @Param bank_account_id path int true "bank account id"
// @Param updateData body entity.UpdateBankAccountInput true "update account info"
// @Success 200 {object} statusResponse "updated bank account"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{organization_id}/bank_accounts/{bank_account_id} [put]
func (h *Handler) updateBankAccount(c *gin.Context) {
	bankAccountId, err := strconv.Atoi(c.Param("bank_account_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid bank_account_id param")
		return
	}
	var updateData entity.UpdateBankAccountInput

	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updateMessage, err := h.services.PutlistAPI.UpdateBankAccount(context.Background(), int64(bankAccountId), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "bank account not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: updateMessage,
	})

}

// @Summary Delete Bank Account
// @Security ApiKeyAuth
// @Tags bank-accounts
// @Description  deletes bank account
// @ID delete-account
// @Accept json
// @Produce json
// @Param organization_id path int true "organization id"
// @Param bank_account_id path int true "bank account id"
// @Success 200 {object} statusResponse "deleted"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{organization_id}/bank_accounts/{bank_account_id} [delete]
func (h *Handler) deleteBankAccount(c *gin.Context) {
	bankAccountId, err := strconv.Atoi(c.Param("bank_account_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid bank_account_id param")
		return
	}

	delMessage, err := h.services.PutlistAPI.DeleteBankAccount(context.Background(), int64(bankAccountId))
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			NewErrorResponse(c, http.StatusNotFound, "bank account not found")
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: delMessage,
	})

}
