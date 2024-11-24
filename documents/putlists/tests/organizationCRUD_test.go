package tests

import (
	"testing"

	docv1 "github.com/DavidG9999/DMS/api/grpc/document_api/gen/go"
	"github.com/DavidG9999/DMS/documents/putlists/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func Test_CRUD_Organization_HappyPath(t *testing.T) {

	ctx, st := suite.NewSuite(t)

	name := gofakeit.Company()
	address := gofakeit.Address().Address
	chief := gofakeit.Name()
	finChief := gofakeit.Name()
	innKpp := randomInnKpp()

	responseCreateOrganization, err := st.OrganizationClient.CreateOrganization(ctx, &docv1.CreateOrganizationRequest{
		Name:     name,
		Address:  address,
		Chief:    chief,
		FinChief: finChief,
		InnKpp:   innKpp,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateOrganization.GetOrganizationId())

	responseGetOrganizations, err := st.OrganizationClient.GetOrganizations(ctx, &docv1.GetOrganizationsRequest{})
	require.NoError(t, err)
	require.NotEmpty(t, responseGetOrganizations.GetOrganizations())

	orgname := gofakeit.Company()
	orgaddress := gofakeit.Address().Address
	orgchief := gofakeit.Name()
	orgfinChief := gofakeit.Name()
	orginnKpp := randomInnKpp()

	t.Run("UpdateOrganizationTest №1", func(t *testing.T) {
		responseUpdateOrganization, err := st.OrganizationClient.UpdateOrganization(ctx, &docv1.UpdateOrganizationRequest{
			OrganizationId: responseCreateOrganization.GetOrganizationId(),
			Name:           &orgname,
			Address:        &orgaddress,
			Chief:          &orgchief,
			FinChief:       &orgfinChief,
			InnKpp:         &orginnKpp,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateOrganization.GetMessage())
		require.Equal(t, responseUpdateOrganization.GetMessage(), "updated organization")
	})

	t.Run("UpdateOrganizationTest №2", func(t *testing.T) {
		responseUpdateOrganization, err := st.OrganizationClient.UpdateOrganization(ctx, &docv1.UpdateOrganizationRequest{
			OrganizationId: responseCreateOrganization.GetOrganizationId(),
			Chief:          &orgchief,
			FinChief:       &orgfinChief,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateOrganization.GetMessage())
		require.Equal(t, responseUpdateOrganization.GetMessage(), "updated organization")
	})

	t.Run("UpdateOrganizationTest №3", func(t *testing.T) {
		responseUpdateOrganization, err := st.OrganizationClient.UpdateOrganization(ctx, &docv1.UpdateOrganizationRequest{
			OrganizationId: responseCreateOrganization.GetOrganizationId(),
			Name:           &orgname,
			Address:        &orgaddress,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateOrganization.GetMessage())
		require.Equal(t, responseUpdateOrganization.GetMessage(), "updated organization")
	})

	delName := gofakeit.Company()
	delAddress := gofakeit.Address().Address
	delChief := gofakeit.Name()
	delFinChief := gofakeit.Name()
	delInnKpp := randomInnKpp()

	responseDeleteCreatedOrganization, err := st.OrganizationClient.CreateOrganization(ctx, &docv1.CreateOrganizationRequest{
		Name:     delName,
		Address:  delAddress,
		Chief:    delChief,
		FinChief: delFinChief,
		InnKpp:   delInnKpp,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteCreatedOrganization.GetOrganizationId())

	responseDeleteOrganization, err := st.OrganizationClient.DeleteOrganization(ctx, &docv1.DeleteOrganizationRequest{
		OrganizationId: responseDeleteCreatedOrganization.GetOrganizationId(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteOrganization.GetMessage())
	require.Equal(t, responseDeleteOrganization.GetMessage(), "deleted")
}

func Test_CRUD_Organization_FailCases(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	name := gofakeit.Company()
	address := gofakeit.Address().Address
	chief := gofakeit.Name()
	finChief := gofakeit.Name()
	innKpp := randomInnKpp()

	t.Run("CreateOrganizationDublicateEmailTest", func(t *testing.T) {
		responseCreateOrganization, err := st.OrganizationClient.CreateOrganization(ctx, &docv1.CreateOrganizationRequest{
			Name:     name,
			Address:  address,
			Chief:    chief,
			FinChief: finChief,
			InnKpp:   innKpp,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseCreateOrganization.GetOrganizationId())

		responseDubleCreateOrganization, err := st.OrganizationClient.CreateOrganization(ctx, &docv1.CreateOrganizationRequest{
			Name:     name,
			Address:  address,
			Chief:    chief,
			FinChief: finChief,
			InnKpp:   innKpp,
		})
		require.Error(t, err)
		require.Empty(t, responseDubleCreateOrganization.GetOrganizationId())
		require.ErrorContains(t, err, "organization already exists")
	})

	createOrganizationTests := []struct {
		name        string
		address     string
		chief       string
		finChief    string
		innKpp      string
		expectedErr string
	}{
		{
			name:        "",
			address:     gofakeit.Address().Address,
			chief:       gofakeit.Name(),
			finChief:    gofakeit.Name(),
			innKpp:      randomInnKpp(),
			expectedErr: "organization name is required",
		},
		{
			address:     gofakeit.Address().Address,
			chief:       gofakeit.Name(),
			finChief:    gofakeit.Name(),
			innKpp:      randomInnKpp(),
			expectedErr: "organization name is required",
		},
		{
			name:        gofakeit.Company(),
			address:     "",
			chief:       gofakeit.Name(),
			finChief:    gofakeit.Name(),
			innKpp:      randomInnKpp(),
			expectedErr: "address is required",
		},
		{
			name:        gofakeit.Company(),
			chief:       gofakeit.Name(),
			finChief:    gofakeit.Name(),
			innKpp:      randomInnKpp(),
			expectedErr: "address is required",
		},
		{
			name:        gofakeit.Company(),
			address:     gofakeit.Address().Address,
			chief:       gofakeit.Name(),
			finChief:    gofakeit.Name(),
			innKpp:      "",
			expectedErr: "inn/kpp is required",
		},
		{
			name:        gofakeit.Company(),
			address:     gofakeit.Address().Address,
			chief:       gofakeit.Name(),
			finChief:    gofakeit.Name(),
			expectedErr: "inn/kpp is required",
		},
		{
			name:        gofakeit.Company(),
			address:     gofakeit.Address().Address,
			chief:       gofakeit.Name(),
			finChief:    gofakeit.Name(),
			innKpp:      invalidInnKppMin(),
			expectedErr: "invalid field format: inn/kpp",
		},
		{
			name:        gofakeit.Company(),
			address:     gofakeit.Address().Address,
			chief:       gofakeit.Name(),
			finChief:    gofakeit.Name(),
			innKpp:      invalidInnKppMax(),
			expectedErr: "invalid field format: inn/kpp",
		},
		{
			name:        gofakeit.Company(),
			address:     gofakeit.Address().Address,
			chief:       "",
			finChief:    gofakeit.Name(),
			innKpp:      randomInnKpp(),
			expectedErr: "chief is required",
		},
		{
			name:        gofakeit.Company(),
			address:     gofakeit.Address().Address,
			finChief:    gofakeit.Name(),
			innKpp:      randomInnKpp(),
			expectedErr: "chief is required",
		},
		{
			name:        gofakeit.Company(),
			address:     gofakeit.Address().Address,
			chief:       gofakeit.Name(),
			finChief:    "",
			innKpp:      randomInnKpp(),
			expectedErr: "financial chief is required",
		},
		{
			name:        gofakeit.Company(),
			address:     gofakeit.Address().Address,
			chief:       gofakeit.Name(),
			innKpp:      randomInnKpp(),
			expectedErr: "financial chief is required",
		},
	}

	for _, test := range createOrganizationTests {
		t.Run("CreateOrganizationFailCases", func(t *testing.T) {
			responseCreateOrganization, err := st.OrganizationClient.CreateOrganization(ctx, &docv1.CreateOrganizationRequest{
				Name:     test.name,
				Address:  test.address,
				Chief:    test.chief,
				FinChief: test.finChief,
				InnKpp:   test.innKpp,
			})
			require.Error(t, err)
			require.Empty(t, responseCreateOrganization.GetOrganizationId())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	beforeUpName := gofakeit.Company()
	beforeUpAddress := gofakeit.Address().Address
	beforeUpChief := gofakeit.Name()
	beforeUpFinChief := gofakeit.Name()
	beforeUpInnKpp := randomInnKpp()

	responseCreateOrganization, err := st.OrganizationClient.CreateOrganization(ctx, &docv1.CreateOrganizationRequest{
		Name:     beforeUpName,
		Address:  beforeUpAddress,
		Chief:    beforeUpChief,
		FinChief: beforeUpFinChief,
		InnKpp:   beforeUpInnKpp,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateOrganization.GetOrganizationId())

	respOrganizationId := responseCreateOrganization.GetOrganizationId()
	updateName := gofakeit.Company()
	updateAddress := gofakeit.Address().Address
	updateInnKpp := randomInnKpp()
	updateChief := gofakeit.Name()
	updateFinChief := gofakeit.Name()
	emptyName := ""
	emptyAddress := ""
	emptyInnKpp := ""
	emptyChief := ""
	emptyFinChief := ""
	notFoundOrganizationId := int64(gofakeit.IntRange(1000, 2000))
	invUpdateInKppMin := invalidInnKppMin()
	invUpdateInKppMax := invalidInnKppMax()

	t.Run("UpdateOrganizationFailCases №1", func(t *testing.T) {
		responseUpdateOrganization, err := st.OrganizationClient.UpdateOrganization(ctx, &docv1.UpdateOrganizationRequest{
			OrganizationId: 0,
			Name:           &updateName,
			Address:        &updateAddress,
			Chief:          &updateChief,
			FinChief:       &updateFinChief,
			InnKpp:         &updateInnKpp,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateOrganization.GetMessage())
		require.ErrorContains(t, err, "organization ID is required")
	})

	t.Run("UpdateOrganizationFailCases №2", func(t *testing.T) {
		responseUpdateOrganization, err := st.OrganizationClient.UpdateOrganization(ctx, &docv1.UpdateOrganizationRequest{
			Name:     &updateName,
			Address:  &updateAddress,
			Chief:    &updateChief,
			FinChief: &updateFinChief,
			InnKpp:   &updateInnKpp,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateOrganization.GetMessage())
		require.ErrorContains(t, err, "organization ID is required")
	})

	t.Run("UpdateOrganizationFailCases №3", func(t *testing.T) {
		responseUpdateOrganization, err := st.OrganizationClient.UpdateOrganization(ctx, &docv1.UpdateOrganizationRequest{
			OrganizationId: notFoundOrganizationId,
			Name:           &updateName,
			Address:        &updateAddress,
			Chief:          &updateChief,
			FinChief:       &updateFinChief,
			InnKpp:         &updateInnKpp,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateOrganization.GetMessage())
		require.ErrorContains(t, err, "organization not found")
	})

	t.Run("UpdateOrganizationFailCases №4", func(t *testing.T) {
		responseUpdateOrganization, err := st.OrganizationClient.UpdateOrganization(ctx, &docv1.UpdateOrganizationRequest{
			OrganizationId: respOrganizationId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateOrganization.GetMessage())
		require.ErrorContains(t, err, "updating structure has no values")
	})

	t.Run("UpdateOrganizationFailCases №5", func(t *testing.T) {
		responseUpdateOrganization, err := st.OrganizationClient.UpdateOrganization(ctx, &docv1.UpdateOrganizationRequest{
			OrganizationId: respOrganizationId,
			Name:           &emptyName,
			Address:        &emptyAddress,
			Chief:          &emptyChief,
			FinChief:       &emptyFinChief,
			InnKpp:         &emptyInnKpp,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateOrganization.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateOrganizationFailCases №6", func(t *testing.T) {
		responseUpdateOrganization, err := st.OrganizationClient.UpdateOrganization(ctx, &docv1.UpdateOrganizationRequest{
			OrganizationId: respOrganizationId,
			InnKpp:         &emptyInnKpp,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateOrganization.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateOrganizationFailCases №7", func(t *testing.T) {
		responseUpdateOrganization, err := st.OrganizationClient.UpdateOrganization(ctx, &docv1.UpdateOrganizationRequest{
			OrganizationId: respOrganizationId,
			Chief:          &emptyChief,
			FinChief:       &emptyFinChief,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateOrganization.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateOrganizationFailCases №8", func(t *testing.T) {
		responseUpdateOrganization, err := st.OrganizationClient.UpdateOrganization(ctx, &docv1.UpdateOrganizationRequest{
			OrganizationId: respOrganizationId,
			Name:           &emptyName,
			Address:        &emptyAddress,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateOrganization.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateOrganizationFailCases №9", func(t *testing.T) {
		responseUpdateOrganization, err := st.OrganizationClient.UpdateOrganization(ctx, &docv1.UpdateOrganizationRequest{
			OrganizationId: respOrganizationId,
			Name:           &updateName,
			Address:        &updateAddress,
			Chief:          &updateChief,
			FinChief:       &updateFinChief,
			InnKpp:         &invUpdateInKppMin,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateOrganization.GetMessage())
		require.ErrorContains(t, err, "invalid field format: inn/kpp")
	})

	t.Run("UpdateOrganizationFailCases №10", func(t *testing.T) {
		responseUpdateOrganization, err := st.OrganizationClient.UpdateOrganization(ctx, &docv1.UpdateOrganizationRequest{
			OrganizationId: respOrganizationId,
			InnKpp:         &invUpdateInKppMax,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateOrganization.GetMessage())
		require.ErrorContains(t, err, "invalid field format: inn/kpp")
	})

	deleteOrganizationTests := []struct {
		organizationID int64
		expectedErr    string
	}{
		{
			organizationID: 0,
			expectedErr:    "organization ID is required",
		},
		{
			expectedErr: "organization ID is required",
		},
		{
			organizationID: int64(gofakeit.IntRange(1000, 2000)),
			expectedErr:    "organization not found",
		},
	}

	for _, test := range deleteOrganizationTests {
		t.Run("DeleteOrganizationFailCases", func(t *testing.T) {
			responseDeleteOrganization, err := st.OrganizationClient.DeleteOrganization(ctx, &docv1.DeleteOrganizationRequest{
				OrganizationId: test.organizationID,
			})
			require.Error(t, err)
			require.Empty(t, responseDeleteOrganization.GetMessage())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

}
