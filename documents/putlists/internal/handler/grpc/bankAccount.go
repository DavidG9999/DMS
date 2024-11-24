package handler

import (
	"context"
	"errors"

	putlistv1 "github.com/DavidG9999/DMS/api/grpc/document_api/putlist/gen/go"
	"github.com/DavidG9999/DMS/documents/putlists/internal/domain/entity"
	"github.com/DavidG9999/DMS/documents/putlists/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BankAccountHandler struct {
	putlistv1.UnimplementedBankAccountServer
	BankAccountService service.BankAccount
}

func (h *BankAccountHandler) CreateBankAccount(ctx context.Context, req *putlistv1.CreateBankAccountRequest) (*putlistv1.CreateBankAccountResponse, error) {
	if err := validateCreateBankAccount(req); err != nil {
		return nil, err
	}

	bankAccount := entity.BankAccount{
		BankAccountNumber: req.GetBankAccountNumber(),
		BankName:          req.GetBankName(),
		BankIdNumber:      req.GetBankIdNumber(),
		OrganizationId:    req.GetOrganizationId(),
	}
	bankAccountId, err := h.BankAccountService.CreateBankAccount(ctx, bankAccount)
	if err != nil {
		if errors.Is(err, service.ErrBankAccExists) {
			return nil, status.Error(codes.AlreadyExists, "bank account already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.CreateBankAccountResponse{
		BankAccountId: bankAccountId,
	}, nil
}

func (h *BankAccountHandler) GetBankAccounts(ctx context.Context, req *putlistv1.GetBankAccountsRequest) (*putlistv1.GetBankAccountsResponse, error) {
	if err := validateGetBankAccounts(req); err != nil {
		return nil, err
	}

	bankAccounts, err := h.BankAccountService.GetBankAccounts(ctx, req.GetOrganizationId())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "bank accounts not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	bankAccountsList := make([]putlistv1.BankAccountEntity, len(bankAccounts))
	for idx, bankAccount := range bankAccounts {
		bankAccountsList[idx].BankAccountId = bankAccount.Id
		bankAccountsList[idx].BankAccountNumber = bankAccount.BankAccountNumber
		bankAccountsList[idx].BankName = bankAccount.BankName
		bankAccountsList[idx].BankIdNumber = bankAccount.BankIdNumber
		bankAccountsList[idx].OrganizationId = bankAccount.OrganizationId
	}

	bankAccountsResp := make([]*putlistv1.BankAccountEntity, 0, len(bankAccounts))
	for id := range bankAccountsList {
		bankAccountsResp = append(bankAccountsResp, &bankAccountsList[id])
	}

	return &putlistv1.GetBankAccountsResponse{
		BankAccounts: bankAccountsResp,
	}, nil
}

func (h *BankAccountHandler) UpdateBankAccount(ctx context.Context, req *putlistv1.UpdateBankAccountRequest) (*putlistv1.UpdateBankAccountResponse, error) {
	if err := validateUpdateBankAccount(req); err != nil {
		return nil, err
	}

	updateData := entity.UpdateBankAccountInput{
		BankAccountNumber: req.BankAccountNumber,
		BankName:          req.BankName,
		BankIdNumber:      req.BankIdNumber,
	}
	err := h.BankAccountService.UpdateBankAccount(ctx, req.GetBankAccountId(), updateData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "bank account not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.UpdateBankAccountResponse{
		Message: "updated bank account",
	}, nil
}

func (h *BankAccountHandler) DeleteBankAccount(ctx context.Context, req *putlistv1.DeleteBankAccountRequest) (*putlistv1.DeleteBankAccountResponse, error) {
	if err := validateDeleteBankAccount(req); err != nil {
		return nil, err
	}

	err := h.BankAccountService.DeleteBankAccount(ctx, req.GetBankAccountId())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "bank account not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &putlistv1.DeleteBankAccountResponse{
		Message: "deleted",
	}, nil
}

func validateCreateBankAccount(req *putlistv1.CreateBankAccountRequest) error {
	if req.GetBankAccountNumber() == "" {
		return status.Error(codes.InvalidArgument, "bank account number is required")
	}
	if req.GetBankName() == "" {
		return status.Error(codes.InvalidArgument, "bank name is required")
	}
	if req.GetBankIdNumber() == "" {
		return status.Error(codes.InvalidArgument, "bank identity number is required")
	}
	if req.GetOrganizationId() == 0 {
		return status.Error(codes.InvalidArgument, "organization ID is required")
	}
	if len(req.GetBankAccountNumber()) != 20 {
		return status.Error(codes.InvalidArgument, "invalid field format: bank account number")
	}
	if len(req.GetBankIdNumber()) != 9 {
		return status.Error(codes.InvalidArgument, "invalid field format: bank identity number")
	}
	return nil
}

func validateGetBankAccounts(req *putlistv1.GetBankAccountsRequest) error {
	if req.GetOrganizationId() == 0 {
		return status.Error(codes.InvalidArgument, "organization ID is required")
	}
	return nil
}

func validateUpdateBankAccount(req *putlistv1.UpdateBankAccountRequest) error {
	if req.GetBankAccountId() == 0 {
		return status.Error(codes.InvalidArgument, "bank account ID is required")
	}
	if req.BankAccountNumber == nil && req.BankName == nil && req.BankIdNumber == nil {
		return status.Error(codes.InvalidArgument, "updating structure has no values")
	}
	if req.BankAccountNumber != nil {
		if *req.BankAccountNumber == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
		if len(*req.BankAccountNumber) != 20 {
			return status.Error(codes.InvalidArgument, "invalid field format: bank account number")
		}
	}
	if req.BankIdNumber != nil {
		if *req.BankIdNumber == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
		if len(*req.BankIdNumber) != 9 {
			return status.Error(codes.InvalidArgument, "invalid field format: bank identity number")
		}
	}
	if req.BankName != nil {
		if *req.BankName == "" {
			return status.Error(codes.InvalidArgument, "updating structure has empty values")
		}
	}
	return nil
}

func validateDeleteBankAccount(req *putlistv1.DeleteBankAccountRequest) error {
	if req.GetBankAccountId() == 0 {
		return status.Error(codes.InvalidArgument, "bank account ID is required")
	}
	return nil
}
