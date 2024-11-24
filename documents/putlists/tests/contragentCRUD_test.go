package tests

import (
	"testing"

	docv1 "github.com/DavidG9999/DMS/api/grpc/document_api/gen/go"
	"github.com/DavidG9999/DMS/documents/putlists/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func Test_CRUD_Contragent_HappyPath(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	name := gofakeit.Company()
	address := gofakeit.Address().Address
	innKpp := randomInnKpp()

	responseCreateContragent, err := st.ContragentClient.CreateContragent(ctx, &docv1.CreateContragentRequest{
		Name:    name,
		Address: address,
		InnKpp:  innKpp,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateContragent.GetContragentId())

	responseGetContragents, err := st.ContragentClient.GetContragents(ctx, &docv1.GetContragentsRequest{})
	require.NoError(t, err)
	require.NotEmpty(t, responseGetContragents.GetContragents())

	nameContragent := gofakeit.Company()
	addressContragent := gofakeit.Address().Address
	innKppContragent := randomInnKpp()
	t.Run("UpdateContragentTest №1", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: responseCreateContragent.GetContragentId(),
			Name:         &nameContragent,
			Address:      &addressContragent,
			InnKpp:       &innKppContragent,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateContragent.GetMessage())
		require.Equal(t, responseUpdateContragent.GetMessage(), "updated contragent")
	})

	t.Run("UpdateContragentTest №2", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: responseCreateContragent.GetContragentId(),
			Name:         &nameContragent,
			Address:      &addressContragent,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateContragent.GetMessage())
		require.Equal(t, responseUpdateContragent.GetMessage(), "updated contragent")
	})

	t.Run("UpdateContragentTest №3", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: responseCreateContragent.GetContragentId(),
			Name:         &nameContragent,
			InnKpp:       &innKppContragent,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateContragent.GetMessage())
		require.Equal(t, responseUpdateContragent.GetMessage(), "updated contragent")
	})

	t.Run("UpdateContragentTest №4", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: responseCreateContragent.GetContragentId(),
			Address:      &addressContragent,
			InnKpp:       &innKppContragent,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateContragent.GetMessage())
		require.Equal(t, responseUpdateContragent.GetMessage(), "updated contragent")
	})

	t.Run("UpdateContragentTest №5", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: responseCreateContragent.GetContragentId(),
			Name:         &nameContragent,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateContragent.GetMessage())
		require.Equal(t, responseUpdateContragent.GetMessage(), "updated contragent")
	})

	t.Run("UpdateContragentTest №6", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: responseCreateContragent.GetContragentId(),
			Address:      &addressContragent,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateContragent.GetMessage())
		require.Equal(t, responseUpdateContragent.GetMessage(), "updated contragent")
	})

	t.Run("UpdateContragentTest №7", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: responseCreateContragent.GetContragentId(),
			InnKpp:       &innKppContragent,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateContragent.GetMessage())
		require.Equal(t, responseUpdateContragent.GetMessage(), "updated contragent")
	})

	delName := gofakeit.Company()
	delAddress := gofakeit.Address().Address
	delInnKpp := randomInnKpp()

	responseDeleteCreatedContragent, err := st.ContragentClient.CreateContragent(ctx, &docv1.CreateContragentRequest{
		Name:    delName,
		Address: delAddress,
		InnKpp:  delInnKpp,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteCreatedContragent.GetContragentId())

	responseDeleteContragent, err := st.ContragentClient.DeleteContragent(ctx, &docv1.DeleteContragentRequest{
		ContragentId: responseDeleteCreatedContragent.GetContragentId(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteContragent.GetMessage())
	require.Equal(t, responseDeleteContragent.GetMessage(), "deleted")
}

func Test_CRUD_Contragent_FailCases(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	name := gofakeit.Company()
	address := gofakeit.Address().Address
	innKpp := randomInnKpp()

	t.Run("CreateContragentDublicateEmailTest", func(t *testing.T) {
		responseCreateContragent, err := st.ContragentClient.CreateContragent(ctx, &docv1.CreateContragentRequest{
			Name:    name,
			Address: address,
			InnKpp:  innKpp,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseCreateContragent.GetContragentId())

		responseDubleCreateContragent, err := st.ContragentClient.CreateContragent(ctx, &docv1.CreateContragentRequest{
			Name:    name,
			Address: address,
			InnKpp:  innKpp,
		})
		require.Error(t, err)
		require.Empty(t, responseDubleCreateContragent.GetContragentId())
		require.ErrorContains(t, err, "contragent already exists")
	})

	createContragentTests := []struct {
		name        string
		address     string
		innKpp      string
		expectedErr string
	}{
		{
			name:        "",
			address:     gofakeit.Address().Address,
			innKpp:      randomInnKpp(),
			expectedErr: "contragent name is required",
		},
		{
			address:     gofakeit.Address().Address,
			innKpp:      randomInnKpp(),
			expectedErr: "contragent name is required",
		},
		{
			name:        gofakeit.Company(),
			address:     "",
			innKpp:      randomInnKpp(),
			expectedErr: "address is required",
		},
		{
			name:        gofakeit.Company(),
			innKpp:      randomInnKpp(),
			expectedErr: "address is required",
		},
		{
			name:        gofakeit.Company(),
			address:     gofakeit.Address().Address,
			innKpp:      "",
			expectedErr: "inn/kpp is required",
		},
		{
			name:        gofakeit.Company(),
			address:     gofakeit.Address().Address,
			expectedErr: "inn/kpp is required",
		},
		{
			name:        gofakeit.Company(),
			address:     gofakeit.Address().Address,
			innKpp:      invalidInnKppMin(),
			expectedErr: "invalid field format: inn/kpp",
		},
		{
			name:        gofakeit.Company(),
			address:     gofakeit.Address().Address,
			innKpp:      invalidInnKppMax(),
			expectedErr: "invalid field format: inn/kpp",
		},
	}

	for _, test := range createContragentTests {
		t.Run("CreateContragentFailCases", func(t *testing.T) {
			responseCreateContragent, err := st.ContragentClient.CreateContragent(ctx, &docv1.CreateContragentRequest{
				Name:    test.name,
				Address: test.address,
				InnKpp:  test.innKpp,
			})
			require.Error(t, err)
			require.Empty(t, responseCreateContragent.GetContragentId())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	beforeUpName := gofakeit.Company()
	beforeUpAddress := gofakeit.Address().Address
	beforeUpInnKpp := randomInnKpp()

	responseCreateContragent, err := st.ContragentClient.CreateContragent(ctx, &docv1.CreateContragentRequest{
		Name:    beforeUpName,
		Address: beforeUpAddress,
		InnKpp:  beforeUpInnKpp,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateContragent.GetContragentId())

	respContragentId := responseCreateContragent.GetContragentId()
	updateName := gofakeit.Company()
	updateAddress := gofakeit.Address().Address
	updateInnKpp := randomInnKpp()
	emptyUpdateName := ""
	emptyUpdateAddress := ""
	emptyUpdateInnKpp := ""
	notFoundContragentId := int64(gofakeit.IntRange(1000, 2000))
	invUpdateInKppMin := invalidInnKppMin()
	invUpdateInKppMax := invalidInnKppMax()

	t.Run("UpdateContragentFailCases №1", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			Name:    &updateName,
			Address: &updateAddress,
			InnKpp:  &updateInnKpp,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateContragent.GetMessage())
		require.ErrorContains(t, err, "contragent ID is required")
	})

	t.Run("UpdateContragentFailCases №2", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: 0,
			Name:         &updateName,
			Address:      &updateAddress,
			InnKpp:       &updateInnKpp,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateContragent.GetMessage())
		require.ErrorContains(t, err, "contragent ID is required")
	})

	t.Run("UpdateContragentFailCases №3", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: notFoundContragentId,
			Name:         &updateName,
			Address:      &updateAddress,
			InnKpp:       &updateInnKpp,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateContragent.GetMessage())
		require.ErrorContains(t, err, "contragent not found")
	})

	t.Run("UpdateContragentFailCases №4", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: respContragentId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateContragent.GetMessage())
		require.ErrorContains(t, err, "updating structure has no values")
	})

	t.Run("UpdateContragentFailCases №5", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: respContragentId,
			Name:         &emptyUpdateName,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateContragent.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateContragentFailCases №6", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: respContragentId,
			Address:      &emptyUpdateAddress,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateContragent.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateContragentFailCases №7", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: respContragentId,
			InnKpp:       &emptyUpdateInnKpp,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateContragent.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateContragentFailCases №8", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: respContragentId,
			Name:         &emptyUpdateName,
			Address:      &emptyUpdateAddress,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateContragent.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateContragentFailCases №9", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: respContragentId,
			Name:         &emptyUpdateName,
			Address:      &emptyUpdateAddress,
			InnKpp:       &emptyUpdateInnKpp,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateContragent.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateContragentFailCases №10", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: respContragentId,
			InnKpp:       &invUpdateInKppMin,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateContragent.GetMessage())
		require.ErrorContains(t, err, "invalid field format: inn/kpp")
	})

	t.Run("UpdateContragentFailCases №11", func(t *testing.T) {
		responseUpdateContragent, err := st.ContragentClient.UpdateContragent(ctx, &docv1.UpdateContragentRequest{
			ContragentId: respContragentId,
			Name:         &updateName,
			Address:      &updateAddress,
			InnKpp:       &invUpdateInKppMax,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateContragent.GetMessage())
		require.ErrorContains(t, err, "invalid field format: inn/kpp")
	})

	deleteContragentTests := []struct {
		contragentID int64
		expectedErr  string
	}{
		{
			contragentID: 0,
			expectedErr:  "contragent ID is required",
		},
		{
			expectedErr: "contragent ID is required",
		},
		{
			contragentID: int64(gofakeit.IntRange(1000, 2000)),
			expectedErr:  "contragent not found",
		},
	}

	for _, test := range deleteContragentTests {
		t.Run("DeleteContragentFailCases", func(t *testing.T) {
			responseDeleteContragent, err := st.ContragentClient.DeleteContragent(ctx, &docv1.DeleteContragentRequest{
				ContragentId: test.contragentID,
			})
			require.Error(t, err)
			require.Empty(t, responseDeleteContragent.GetMessage())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

}
