package tests

import (
	"testing"

	docv1 "github.com/DavidG9999/DMS/api/grpc/document_api/gen/go"
	"github.com/DavidG9999/DMS/documents/putlists/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func Test_CRUD_Dispetcher_HappyPath(t *testing.T) {

	ctx, st := suite.NewSuite(t)

	fullname := gofakeit.Name()

	responseCreateDispetcher, err := st.DispetcherClient.CreateDispetcher(ctx, &docv1.CreateDispetcherRequest{
		FullName: fullname,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateDispetcher.GetDispetcherId())

	responseGetDispetchers, err := st.DispetcherClient.GetDispetchers(ctx, &docv1.GetDispetchersRequest{})
	require.NoError(t, err)
	require.NotEmpty(t, responseGetDispetchers.GetDispetchers())

	dispName := gofakeit.Name()
	t.Run("UpdateDispetcherTest", func(t *testing.T) {
		responseUpdateDispetcher, err := st.DispetcherClient.UpdateDispetcher(ctx, &docv1.UpdateDispetcherRequest{
			DispetcherId: responseCreateDispetcher.GetDispetcherId(),
			FullName:     &dispName,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateDispetcher.GetMessage())
		require.Equal(t, responseUpdateDispetcher.GetMessage(), "updated dispetcher")
	})

	delFullName := gofakeit.Name()

	responseDeleteCreatedDispetcher, err := st.DispetcherClient.CreateDispetcher(ctx, &docv1.CreateDispetcherRequest{
		FullName: delFullName,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteCreatedDispetcher.GetDispetcherId())

	responseDeleteDispetcher, err := st.DispetcherClient.DeleteDispetcher(ctx, &docv1.DeleteDispetcherRequest{
		DispetcherId: responseDeleteCreatedDispetcher.GetDispetcherId(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteDispetcher.GetMessage())
	require.Equal(t, responseDeleteDispetcher.GetMessage(), "deleted")
}

func Test_CRUD_Dispetcher_FailCases(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	createDispetcherTests := []struct {
		fullname    string
		expectedErr string
	}{
		{
			fullname:    "",
			expectedErr: "fullname is required",
		},
		{
			expectedErr: "fullname is required",
		},
	}

	for _, test := range createDispetcherTests {
		t.Run("CreateDispetcherFailCases", func(t *testing.T) {
			responseCreateDispetcher, err := st.DispetcherClient.CreateDispetcher(ctx, &docv1.CreateDispetcherRequest{
				FullName: test.fullname,
			})
			require.Error(t, err)
			require.Empty(t, responseCreateDispetcher.GetDispetcherId())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	beforeUpFullName := gofakeit.Name()

	responseCreateDispetcher, err := st.DispetcherClient.CreateDispetcher(ctx, &docv1.CreateDispetcherRequest{
		FullName: beforeUpFullName,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateDispetcher.GetDispetcherId())

	respDispetcherId := responseCreateDispetcher.GetDispetcherId()
	updateFullName := gofakeit.Name()
	emptyFullName := ""
	notFountDisptcherId := int64(gofakeit.IntRange(1000, 2000))

	t.Run("UpdateDispetcherFailCases №1", func(t *testing.T) {
		responseUpdateDispetcher, err := st.DispetcherClient.UpdateDispetcher(ctx, &docv1.UpdateDispetcherRequest{
			DispetcherId: 0,
			FullName:     &updateFullName,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDispetcher.GetMessage())
		require.ErrorContains(t, err, "dispetcher ID is required")
	})

	t.Run("UpdateDispetcherFailCases №2", func(t *testing.T) {
		responseUpdateDispetcher, err := st.DispetcherClient.UpdateDispetcher(ctx, &docv1.UpdateDispetcherRequest{
			FullName: &updateFullName,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDispetcher.GetMessage())
		require.ErrorContains(t, err, "dispetcher ID is required")
	})

	t.Run("UpdateDispetcherFailCases №3", func(t *testing.T) {
		responseUpdateDispetcher, err := st.DispetcherClient.UpdateDispetcher(ctx, &docv1.UpdateDispetcherRequest{
			DispetcherId: notFountDisptcherId,
			FullName:     &updateFullName,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDispetcher.GetMessage())
		require.ErrorContains(t, err, "dispetcher not found")
	})

	t.Run("UpdateDispetcherFailCases №4", func(t *testing.T) {
		responseUpdateDispetcher, err := st.DispetcherClient.UpdateDispetcher(ctx, &docv1.UpdateDispetcherRequest{
			DispetcherId: respDispetcherId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDispetcher.GetMessage())
		require.ErrorContains(t, err, "updating structure has no values")
	})

	t.Run("UpdateDispetcherFailCases №5", func(t *testing.T) {
		responseUpdateDispetcher, err := st.DispetcherClient.UpdateDispetcher(ctx, &docv1.UpdateDispetcherRequest{
			DispetcherId: respDispetcherId,
			FullName:     &emptyFullName,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDispetcher.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	deleteDispetcherTests := []struct {
		dispetcherID int64
		expectedErr  string
	}{
		{
			dispetcherID: 0,
			expectedErr:  "dispetcher ID is required",
		},
		{
			expectedErr: "dispetcher ID is required",
		},
		{
			dispetcherID: int64(gofakeit.IntRange(1000, 2000)),
			expectedErr:  "dispetcher not found",
		},
	}

	for _, test := range deleteDispetcherTests {
		t.Run("DeleteDispetcherFailCases", func(t *testing.T) {
			responseDeleteDispetcher, err := st.DispetcherClient.DeleteDispetcher(ctx, &docv1.DeleteDispetcherRequest{
				DispetcherId: test.dispetcherID,
			})
			require.Error(t, err)
			require.Empty(t, responseDeleteDispetcher.GetMessage())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

}
