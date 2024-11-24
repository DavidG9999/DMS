package tests

import (
	"testing"

	docv1 "github.com/DavidG9999/DMS/api/grpc/document_api/gen/go"
	"github.com/DavidG9999/DMS/documents/putlists/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

const organizationId int64 = 1

func Test_CRUD_BankAccount_HappyPath(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	randCreateBank := randomBank()
	bankAccNum := randomBankAccountNumber()
	bankName := randCreateBank.bankName
	bankIdNum := randCreateBank.bankIdNubmer

	responseCreateBankAccount, err := st.BankAccountClient.CreateBankAccount(ctx, &docv1.CreateBankAccountRequest{
		BankAccountNumber: bankAccNum,
		BankName:          bankName,
		BankIdNumber:      bankIdNum,
		OrganizationId:    organizationId,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateBankAccount.GetBankAccountId())

	responseGetBankAccounts, err := st.BankAccountClient.GetBankAccounts(ctx, &docv1.GetBankAccountsRequest{
		OrganizationId: organizationId,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseGetBankAccounts.GetBankAccounts())

	randUpdateBank := randomBank()

	t.Run("UpdateBankAccountTest №1", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId:     responseCreateBankAccount.GetBankAccountId(),
			BankAccountNumber: updateRandomBankAccountNumber(),
			BankName:          &randUpdateBank.bankName,
			BankIdNumber:      &randUpdateBank.bankIdNubmer,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateBankAccount.GetMessage())
		require.Equal(t, responseUpdateBankAccount.GetMessage(), "updated bank account")
	})

	t.Run("UpdateBankAccountTest №2", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId:     responseCreateBankAccount.GetBankAccountId(),
			BankAccountNumber: updateRandomBankAccountNumber(),
			BankName:          &randUpdateBank.bankName,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateBankAccount.GetMessage())
		require.Equal(t, responseUpdateBankAccount.GetMessage(), "updated bank account")
	})

	t.Run("UpdateBankAccountTest №3", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId:     responseCreateBankAccount.GetBankAccountId(),
			BankAccountNumber: updateRandomBankAccountNumber(),
			BankIdNumber:      &randUpdateBank.bankIdNubmer,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateBankAccount.GetMessage())
		require.Equal(t, responseUpdateBankAccount.GetMessage(), "updated bank account")
	})

	t.Run("UpdateBankAccountTest №4", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId: responseCreateBankAccount.GetBankAccountId(),
			BankName:      &randUpdateBank.bankName,
			BankIdNumber:  &randUpdateBank.bankIdNubmer,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateBankAccount.GetMessage())
		require.Equal(t, responseUpdateBankAccount.GetMessage(), "updated bank account")
	})

	t.Run("UpdateBankAccountTest №5", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId:     responseCreateBankAccount.GetBankAccountId(),
			BankAccountNumber: updateRandomBankAccountNumber(),
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateBankAccount.GetMessage())
		require.Equal(t, responseUpdateBankAccount.GetMessage(), "updated bank account")
	})

	t.Run("UpdateBankAccountTest №6", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId: responseCreateBankAccount.GetBankAccountId(),
			BankIdNumber:  &randUpdateBank.bankIdNubmer,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateBankAccount.GetMessage())
		require.Equal(t, responseUpdateBankAccount.GetMessage(), "updated bank account")
	})

	t.Run("UpdateBankAccountTest №7", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId: responseCreateBankAccount.GetBankAccountId(),
			BankName:      &randUpdateBank.bankName,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateBankAccount.GetMessage())
		require.Equal(t, responseUpdateBankAccount.GetMessage(), "updated bank account")
	})

	delrandCreateBank := randomBank()
	delbankAccNum := randomBankAccountNumber()
	delbankName := delrandCreateBank.bankName
	delbankIdNum := delrandCreateBank.bankIdNubmer

	responseDeleteCreatedBankAccount, err := st.BankAccountClient.CreateBankAccount(ctx, &docv1.CreateBankAccountRequest{
		BankAccountNumber: delbankAccNum,
		BankName:          delbankName,
		BankIdNumber:      delbankIdNum,
		OrganizationId:    organizationId,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteCreatedBankAccount.GetBankAccountId())

	responseDeleteBankAccount, err := st.BankAccountClient.DeleteBankAccount(ctx, &docv1.DeleteBankAccountRequest{
		BankAccountId: responseDeleteCreatedBankAccount.GetBankAccountId(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteBankAccount.GetMessage())
	require.Equal(t, responseDeleteBankAccount.GetMessage(), "deleted")
}

func Test_CRUD_BankAccount_FailCases(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	randCreateBank := randomBank()
	bankAccNum := randomBankAccountNumber()
	bankName := randCreateBank.bankName
	bankIdNum := randCreateBank.bankIdNubmer

	t.Run("CreateBankAccountDublicateEmailTest", func(t *testing.T) {
		responseCreateBankAccount, err := st.BankAccountClient.CreateBankAccount(ctx, &docv1.CreateBankAccountRequest{
			BankAccountNumber: bankAccNum,
			BankName:          bankName,
			BankIdNumber:      bankIdNum,
			OrganizationId:    organizationId,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseCreateBankAccount.GetBankAccountId())

		responseDubleCreateBankAccount, err := st.BankAccountClient.CreateBankAccount(ctx, &docv1.CreateBankAccountRequest{
			BankAccountNumber: bankAccNum,
			BankName:          bankName,
			BankIdNumber:      bankIdNum,
			OrganizationId:    organizationId,
		})
		require.Error(t, err)
		require.Empty(t, responseDubleCreateBankAccount.GetBankAccountId())
		require.ErrorContains(t, err, "bank account already exists")
	})

	createBankAccountTests := []struct {
		bankAccNum     string
		bankName       string
		bankIdNum      string
		organizationId int64
		expectedErr    string
	}{
		{
			bankAccNum:     "",
			bankName:       randCreateBank.bankName,
			bankIdNum:      randCreateBank.bankIdNubmer,
			organizationId: organizationId,
			expectedErr:    "bank account number is required",
		},
		{
			bankAccNum:     randomBankAccountNumber(),
			bankName:       "",
			bankIdNum:      randCreateBank.bankIdNubmer,
			organizationId: organizationId,
			expectedErr:    "bank name is required",
		},
		{
			bankAccNum:     randomBankAccountNumber(),
			bankName:       randCreateBank.bankName,
			bankIdNum:      "",
			organizationId: organizationId,
			expectedErr:    "bank identity number is required",
		},
		{
			bankAccNum:     randomBankAccountNumber(),
			bankName:       randCreateBank.bankName,
			bankIdNum:      randCreateBank.bankIdNubmer,
			organizationId: 0,
			expectedErr:    "organization ID is required",
		},
		{
			bankAccNum:     invalidBankAccountNumberMin(),
			bankName:       randCreateBank.bankName,
			bankIdNum:      randCreateBank.bankIdNubmer,
			organizationId: organizationId,
			expectedErr:    "invalid field format: bank account number",
		},
		{
			bankAccNum:     invalidBankAccountNumberMax(),
			bankName:       randCreateBank.bankName,
			bankIdNum:      randCreateBank.bankIdNubmer,
			organizationId: organizationId,
			expectedErr:    "invalid field format: bank account number",
		},
		{
			bankAccNum:     randomBankAccountNumber(),
			bankName:       randCreateBank.bankName,
			bankIdNum:      "12345678",
			organizationId: organizationId,
			expectedErr:    "invalid field format: bank identity number",
		},
		{
			bankAccNum:     randomBankAccountNumber(),
			bankName:       randCreateBank.bankName,
			bankIdNum:      "1234567890",
			organizationId: organizationId,
			expectedErr:    "invalid field format: bank identity number",
		},
	}

	for _, test := range createBankAccountTests {
		t.Run("CreateBankAccountFailCases", func(t *testing.T) {
			responseCreateBankAccount, err := st.BankAccountClient.CreateBankAccount(ctx, &docv1.CreateBankAccountRequest{
				BankAccountNumber: test.bankAccNum,
				BankName:          test.bankName,
				BankIdNumber:      test.bankIdNum,
				OrganizationId:    test.organizationId,
			})
			require.Error(t, err)
			require.Empty(t, responseCreateBankAccount.GetBankAccountId())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	randUpBank := randomBank()
	beforeUpBankAccNum := randomBankAccountNumber()
	beforeUpBankName := randUpBank.bankName
	beforeUpBankIdNum := randUpBank.bankIdNubmer
	beforeUpOrganizationId := organizationId

	responseCreateBankAccount, err := st.BankAccountClient.CreateBankAccount(ctx, &docv1.CreateBankAccountRequest{
		BankAccountNumber: beforeUpBankAccNum,
		BankName:          beforeUpBankName,
		BankIdNumber:      beforeUpBankIdNum,
		OrganizationId:    beforeUpOrganizationId,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateBankAccount.GetBankAccountId())

	respBankAccId := responseCreateBankAccount.GetBankAccountId()
	randBankAccNum := randomBankAccountNumber()
	randBankName := randCreateBank.bankName
	randBankIdNum := randCreateBank.bankIdNubmer
	notFoundBankAccId := int64(gofakeit.IntRange(1000, 2000))
	emptyBankAccNum := ""
	emptyBankName := ""
	emptyBankIdNum := ""
	invBankAccNumMin := invalidBankAccountNumberMin()
	invBankAccNumMax := invalidBankAccountNumberMax()
	invBankIdNumMin := "12345678"
	invBankIdNumMax := "1234567890"

	t.Run("UpdateBankAccountFailCases №1", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountNumber: &randBankAccNum,
			BankName:          &randBankName,
			BankIdNumber:      &randBankIdNum,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateBankAccount.GetMessage())
		require.ErrorContains(t, err, "bank account ID is required")
	})

	t.Run("UpdateBankAccountFailCases №2", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId:     notFoundBankAccId,
			BankAccountNumber: &bankAccNum,
			BankName:          &bankName,
			BankIdNumber:      &bankIdNum,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateBankAccount.GetMessage())
		require.ErrorContains(t, err, "bank account not found")
	})

	t.Run("UpdateBankAccountFailCases №3", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId: respBankAccId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateBankAccount.GetMessage())
		require.ErrorContains(t, err, "updating structure has no values")
	})

	t.Run("UpdateBankAccountFailCases №4", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId:     respBankAccId,
			BankAccountNumber: &emptyBankAccNum,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateBankAccount.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateBankAccountFailCases №5", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId: respBankAccId,
			BankName:      &emptyBankName,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateBankAccount.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateBankAccountFailCases №6", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId: respBankAccId,
			BankIdNumber:  &emptyBankIdNum,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateBankAccount.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateBankAccountFailCases №7", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId:     respBankAccId,
			BankAccountNumber: &emptyBankAccNum,
			BankName:          &emptyBankName,
			BankIdNumber:      &emptyBankIdNum,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateBankAccount.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateBankAccountFailCases №8", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId:     respBankAccId,
			BankAccountNumber: &invBankAccNumMin,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateBankAccount.GetMessage())
		require.ErrorContains(t, err, "invalid field format: bank account number")
	})

	t.Run("UpdateBankAccountFailCases №9", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId:     respBankAccId,
			BankAccountNumber: &invBankAccNumMax,
			BankName:          &bankName,
			BankIdNumber:      &bankIdNum,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateBankAccount.GetMessage())
		require.ErrorContains(t, err, "invalid field format: bank account number")
	})

	t.Run("UpdateBankAccountFailCases №10", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId: respBankAccId,
			BankName:      &bankName,
			BankIdNumber:  &invBankIdNumMin,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateBankAccount.GetMessage())
		require.ErrorContains(t, err, "invalid field format: bank identity number")
	})

	t.Run("UpdateBankAccountFailCases №11", func(t *testing.T) {
		responseUpdateBankAccount, err := st.BankAccountClient.UpdateBankAccount(ctx, &docv1.UpdateBankAccountRequest{
			BankAccountId: respBankAccId,
			BankIdNumber:  &invBankIdNumMax,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateBankAccount.GetMessage())
		require.ErrorContains(t, err, "invalid field format: bank identity number")
	})

	deleteBankAccountTests := []struct {
		bankaccountID int64
		expectedErr   string
	}{
		{
			bankaccountID: 0,
			expectedErr:   "bank account ID is required",
		},
		{
			expectedErr: "bank account ID is required",
		},
		{
			bankaccountID: int64(gofakeit.IntRange(1000, 2000)),
			expectedErr:   "bank account not found",
		},
	}

	for _, test := range deleteBankAccountTests {
		t.Run("DeleteBankAccountFailCases", func(t *testing.T) {
			responseDeleteBankAccount, err := st.BankAccountClient.DeleteBankAccount(ctx, &docv1.DeleteBankAccountRequest{
				BankAccountId: test.bankaccountID,
			})
			require.Error(t, err)
			require.Empty(t, responseDeleteBankAccount.GetMessage())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

}
