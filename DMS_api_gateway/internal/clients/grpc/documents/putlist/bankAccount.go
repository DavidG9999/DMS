package putlistgrpc

import (
	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	"context"
	"fmt"

	putlistv1 "github.com/DavidG9999/DMS/api/grpc/document_api/putlist/gen/go"
)

func (pc *PutlistClient) CreateBankAccount(ctx context.Context, organizationId int64, bankAccNum, bankName, bankIdNum string) (int64, error) {
	const op = "grpc.ClientCreateBankAccount"

	resp, err := pc.apiBankAccount.CreateBankAccount(ctx, &putlistv1.CreateBankAccountRequest{BankAccountNumber: bankAccNum, BankName: bankName, BankIdNumber: bankIdNum, OrganizationId: organizationId})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetBankAccountId(), nil
}

func (pc *PutlistClient) GetBankAccounts(ctx context.Context,  organizationId int64) ([]entity.BankAccount, error) {
	const op = "grpc.ClientGetBankAccounts"

	resp, err := pc.apiBankAccount.GetBankAccounts(ctx, &putlistv1.GetBankAccountsRequest{OrganizationId: organizationId})
	if err != nil {
		return []entity.BankAccount{}, fmt.Errorf("%s: %w", op, err)
	}

	bankAccs := make([]entity.BankAccount, len(resp.BankAccounts))
	for id, bankAcc := range resp.GetBankAccounts() {
		bankAccs[id].Id = bankAcc.BankAccountId
		bankAccs[id].BankAccountNumber = bankAcc.BankAccountNumber
		bankAccs[id].BankName = bankAcc.BankName
		bankAccs[id].BankIdNumber = bankAcc.BankIdNumber
	}
	return bankAccs, nil
}

func (pc *PutlistClient) UpdateBankAccount(ctx context.Context, bankAccId int64, bankAccNum, bankName, bankIdNum *string) (string, error) {
	const op = "grpc.ClientUpdateBankAccount"

	resp, err := pc.apiBankAccount.UpdateBankAccount(ctx, &putlistv1.UpdateBankAccountRequest{BankAccountId: bankAccId, BankAccountNumber: bankAccNum, BankName: bankName, BankIdNumber: bankIdNum})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}

func (pc *PutlistClient) DeleteBankAccount(ctx context.Context, bankAccId int64) (string, error) {
	const op = "grpc.ClientDeleteBankAccount"

	resp, err := pc.apiBankAccount.DeleteBankAccount(ctx, &putlistv1.DeleteBankAccountRequest{BankAccountId: bankAccId})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}
