package tests

import (
	"testing"

	docv1 "github.com/DavidG9999/DMS/api/grpc/document_api/gen/go"
	"github.com/DavidG9999/DMS/documents/putlists/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func Test_CRUD_Auto_HappyPath(t *testing.T) {

	ctx, st := suite.NewSuite(t)

	randCreateAuto := randomTruck()
	brand := randCreateAuto.brand
	model := randCreateAuto.model
	stateNumber := randomStateNumber()

	responseCreateAuto, err := st.AutoClient.CreateAuto(ctx, &docv1.CreateAutoRequest{
		Brand:       brand,
		Model:       model,
		StateNumber: stateNumber,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateAuto.GetAutoId())

	responseGetAutos, err := st.AutoClient.GetAutos(ctx, &docv1.GetAutosRequest{})
	require.NoError(t, err)
	require.NotEmpty(t, responseGetAutos.GetAutos())

	randUpdateAuto := randomTruck()

	t.Run("UpdateAutoTest №1", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId:      responseCreateAuto.GetAutoId(),
			Brand:       &randUpdateAuto.brand,
			Model:       &randUpdateAuto.model,
			StateNumber: updateRandomStateNumber(),
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateAuto.GetMessage())
		require.Equal(t, responseUpdateAuto.GetMessage(), "updated auto")
	})

	t.Run("UpdateAutoTest №2", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId:      responseCreateAuto.GetAutoId(),
			Model:       &randUpdateAuto.model,
			StateNumber: updateRandomStateNumber(),
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateAuto.GetMessage())
		require.Equal(t, responseUpdateAuto.GetMessage(), "updated auto")
	})

	t.Run("UpdateAutoTest №3", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId:      responseCreateAuto.GetAutoId(),
			Brand:       &randUpdateAuto.brand,
			StateNumber: updateRandomStateNumber(),
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateAuto.GetMessage())
		require.Equal(t, responseUpdateAuto.GetMessage(), "updated auto")
	})

	t.Run("UpdateAutoTest №4", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId: responseCreateAuto.GetAutoId(),
			Brand:  &randUpdateAuto.brand,
			Model:  &randUpdateAuto.model,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateAuto.GetMessage())
		require.Equal(t, responseUpdateAuto.GetMessage(), "updated auto")
	})

	t.Run("UpdateAutoTest №5", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId: responseCreateAuto.GetAutoId(),
			Brand:  &randUpdateAuto.brand,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateAuto.GetMessage())
		require.Equal(t, responseUpdateAuto.GetMessage(), "updated auto")
	})

	t.Run("UpdateAutoTest №6", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId: responseCreateAuto.GetAutoId(),
			Model:  &randUpdateAuto.model,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateAuto.GetMessage())
		require.Equal(t, responseUpdateAuto.GetMessage(), "updated auto")
	})

	t.Run("UpdateAutoTest №7", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId:      responseCreateAuto.GetAutoId(),
			StateNumber: updateRandomStateNumber(),
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateAuto.GetMessage())
		require.Equal(t, responseUpdateAuto.GetMessage(), "updated auto")
	})

	randDeleteAuto := randomTruck()
	delBrand := randDeleteAuto.brand
	delModel := randDeleteAuto.model
	delStateNumber := randomStateNumber()

	responseDeleteCreatedAuto, err := st.AutoClient.CreateAuto(ctx, &docv1.CreateAutoRequest{
		Brand:       delBrand,
		Model:       delModel,
		StateNumber: delStateNumber,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteCreatedAuto.GetAutoId())

	responseDeleteAuto, err := st.AutoClient.DeleteAuto(ctx, &docv1.DeleteAutoRequest{
		AutoId: responseDeleteCreatedAuto.GetAutoId(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteAuto.GetMessage())
	require.Equal(t, responseDeleteAuto.GetMessage(), "deleted")
}

func Test_CRUD_Auto_FailCases(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	randCreateAuto := randomTruck()
	brand := randCreateAuto.brand
	model := randCreateAuto.model
	stateNumber := randomStateNumber()

	t.Run("CreateAutoDublicateEmailTest", func(t *testing.T) {
		responseCreateAuto, err := st.AutoClient.CreateAuto(ctx, &docv1.CreateAutoRequest{
			Brand:       brand,
			Model:       model,
			StateNumber: stateNumber,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseCreateAuto.GetAutoId())

		responseDubleCreateAuto, err := st.AutoClient.CreateAuto(ctx, &docv1.CreateAutoRequest{
			Brand:       brand,
			Model:       model,
			StateNumber: stateNumber,
		})
		require.Error(t, err)
		require.Empty(t, responseDubleCreateAuto.GetAutoId())
		require.ErrorContains(t, err, "auto already exists")
	})

	createAutoTests := []struct {
		brand       string
		model       string
		stateNumber string
		expectedErr string
	}{
		{
			brand:       "",
			model:       randCreateAuto.model,
			stateNumber: randomStateNumber(),
			expectedErr: "brand is required",
		},
		{
			model:       randCreateAuto.model,
			stateNumber: randomStateNumber(),
			expectedErr: "brand is required",
		},
		{
			brand:       randCreateAuto.brand,
			model:       "",
			stateNumber: randomStateNumber(),
			expectedErr: "model is required",
		},
		{
			brand:       randCreateAuto.brand,
			stateNumber: randomStateNumber(),
			expectedErr: "model is required",
		},
		{
			brand:       randCreateAuto.brand,
			model:       randCreateAuto.model,
			stateNumber: "",
			expectedErr: "state number is required",
		},
		{
			brand:       randCreateAuto.brand,
			model:       randCreateAuto.model,
			expectedErr: "state number is required",
		},
		{
			brand:       randCreateAuto.brand,
			model:       randCreateAuto.model,
			stateNumber: invalidStateNumberMin(),
			expectedErr: "invalid field format: state number",
		},
		{
			brand:       randCreateAuto.brand,
			model:       randCreateAuto.model,
			stateNumber: invalidStateNumberMax(),
			expectedErr: "invalid field format: state number",
		},
	}

	for _, test := range createAutoTests {
		t.Run("CreateAutoFailCases", func(t *testing.T) {
			responseCreateAuto, err := st.AutoClient.CreateAuto(ctx, &docv1.CreateAutoRequest{
				Brand:       test.brand,
				Model:       test.model,
				StateNumber: test.stateNumber,
			})
			require.Error(t, err)
			require.Empty(t, responseCreateAuto.GetAutoId())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	randBeforeUpdateAuto := randomTruck()
	beforeUpBrand := randBeforeUpdateAuto.brand
	beforeUpModel := randBeforeUpdateAuto.model
	beforeUpStateNumber := randomStateNumber()

	responseCreateAuto, err := st.AutoClient.CreateAuto(ctx, &docv1.CreateAutoRequest{
		Brand:       beforeUpBrand,
		Model:       beforeUpModel,
		StateNumber: beforeUpStateNumber,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateAuto.GetAutoId())

	randUpdateAuto := randomTruck()
	randStateNumber := randomStateNumber()
	notFoundAutoID := int64(gofakeit.IntRange(1000, 2000))
	emptyBrand := ""
	emptyModel := ""
	emptyStateNumber := ""
	invStateNumberMin := invalidStateNumberMin()
	invStateNumberMax := invalidStateNumberMax()

	t.Run("UpdateAutoFailCases №1", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId:      0,
			Brand:       &randUpdateAuto.brand,
			Model:       &randUpdateAuto.model,
			StateNumber: &randStateNumber,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateAuto.GetMessage())
		require.ErrorContains(t, err, "auto ID is required")
	})

	t.Run("UpdateAutoFailCases №2", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			Brand:       &randUpdateAuto.brand,
			Model:       &randUpdateAuto.model,
			StateNumber: &randStateNumber,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateAuto.GetMessage())
		require.ErrorContains(t, err, "auto ID is required")
	})

	t.Run("UpdateAutoFailCases №3", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId:      notFoundAutoID,
			Brand:       &randUpdateAuto.brand,
			Model:       &randUpdateAuto.model,
			StateNumber: &randStateNumber,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateAuto.GetMessage())
		require.ErrorContains(t, err, "auto not found")
	})

	t.Run("UpdateAutoFailCases №4", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId: responseCreateAuto.GetAutoId(),
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateAuto.GetMessage())
		require.ErrorContains(t, err, "updating structure has no values")
	})

	t.Run("UpdateAutoFailCases №5", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId:      responseCreateAuto.GetAutoId(),
			Brand:       &emptyBrand,
			Model:       &emptyModel,
			StateNumber: &randStateNumber,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateAuto.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateAutoFailCases №6", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId: responseCreateAuto.GetAutoId(),
			Brand:  &emptyBrand,
			Model:  &emptyModel,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateAuto.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateAutoFailCases №7", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId:      responseCreateAuto.GetAutoId(),
			StateNumber: &invStateNumberMin,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateAuto.GetMessage())
		require.ErrorContains(t, err, "invalid field format: state number")
	})

	t.Run("UpdateAutoFailCases №8", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId:      responseCreateAuto.GetAutoId(),
			Brand:       &randUpdateAuto.brand,
			Model:       &randUpdateAuto.model,
			StateNumber: &invStateNumberMax,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateAuto.GetMessage())
		require.ErrorContains(t, err, "invalid field format: state number")
	})

	t.Run("UpdateAutoFailCases №9", func(t *testing.T) {
		responseUpdateAuto, err := st.AutoClient.UpdateAuto(ctx, &docv1.UpdateAutoRequest{
			AutoId:      responseCreateAuto.GetAutoId(),
			StateNumber: &emptyStateNumber,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateAuto.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	deleteAutoTests := []struct {
		autoID      int64
		expectedErr string
	}{
		{
			autoID:      0,
			expectedErr: "auto ID is required",
		},
		{
			expectedErr: "auto ID is required",
		},
		{
			autoID:      int64(gofakeit.IntRange(1000, 2000)),
			expectedErr: "auto not found",
		},
	}

	for _, test := range deleteAutoTests {
		t.Run("DeleteAutoFailCases", func(t *testing.T) {
			responseDeleteAuto, err := st.AutoClient.DeleteAuto(ctx, &docv1.DeleteAutoRequest{
				AutoId: test.autoID,
			})
			require.Error(t, err)
			require.Empty(t, responseDeleteAuto.GetMessage())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

}
