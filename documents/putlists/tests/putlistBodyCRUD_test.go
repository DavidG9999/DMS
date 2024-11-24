package tests

import (
	"math/rand/v2"
	"testing"

	docv1 "github.com/DavidG9999/DMS/api/grpc/document_api/gen/go"
	"github.com/DavidG9999/DMS/documents/putlists/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

const (
	putlistNumber      int64 = 1
	emptyPutlistNumber int64 = 0
	contragentId       int64 = 1
	emptyContragentId  int64 = 0
)

func Test_CRUD_PutlistBody_HappyPath(t *testing.T) {
	ctx, st := suite.NewSuite(t)
	timeWith := "2024-11-13 15:00:00"
	timeFor := "2024-11-13 17:00:00"
	number := int64(rand.IntN(100))
	item := "test"

	responseCreatePutlistBody, err := st.PutlistBodyClient.CreatePutlistBody(ctx, &docv1.CreatePutlistBodyRequest{
		PutlistNumber: putlistNumber,
		Number:        number,
		ContragentId:  contragentId,
		Item:          item,
		TimeWith:      timeWith,
		TimeFor:       timeFor,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreatePutlistBody.GetPutlistBodyId())

	responseGetPutlistBodies, err := st.PutlistBodyClient.GetPutlistBodies(ctx, &docv1.GetPutlistBodiesRequest{
		PutlistNumber: putlistNumber,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseGetPutlistBodies.GetPutlistBodies())

	putBodynumber := number
	putBodycontragentId := contragentId
	putBodyitem := item
	putBodytimeWith := timeWith
	putBodytimeFor := timeFor
	t.Run("UpdatePutlistBodyTest №1", func(t *testing.T) {
		responseUpdatePutlistBody, err := st.PutlistBodyClient.UpdatePutlistBody(ctx, &docv1.UpdatePutlistBodyRequest{
			PutlistBodyId: responseCreatePutlistBody.GetPutlistBodyId(),
			Number:        &putBodynumber,
			ContragentId:  &putBodycontragentId,
			Item:          &putBodyitem,
			TimeWith:      &putBodytimeWith,
			TimeFor:       &putBodytimeFor,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdatePutlistBody.GetMessage())
		require.Equal(t, responseUpdatePutlistBody.GetMessage(), "updated putlist body")
	})

	t.Run("UpdatePutlistBodyTest №2", func(t *testing.T) {
		responseUpdatePutlistBody, err := st.PutlistBodyClient.UpdatePutlistBody(ctx, &docv1.UpdatePutlistBodyRequest{
			PutlistBodyId: responseCreatePutlistBody.GetPutlistBodyId(),
			Number:        &putBodynumber,
			ContragentId:  &putBodycontragentId,
			Item:          &putBodyitem,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdatePutlistBody.GetMessage())
		require.Equal(t, responseUpdatePutlistBody.GetMessage(), "updated putlist body")
	})

	t.Run("UpdatePutlistBodyTest №3", func(t *testing.T) {
		responseUpdatePutlistBody, err := st.PutlistBodyClient.UpdatePutlistBody(ctx, &docv1.UpdatePutlistBodyRequest{
			PutlistBodyId: responseCreatePutlistBody.GetPutlistBodyId(),
			TimeWith:      &putBodytimeWith,
			TimeFor:       &putBodytimeFor,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdatePutlistBody.GetMessage())
		require.Equal(t, responseUpdatePutlistBody.GetMessage(), "updated putlist body")
	})

	delparseTimeWith := "2024-11-13 15:00:00"
	delparseTimeFor := "2024-11-13 17:00:00"
	delnumber := number
	delitem := gofakeit.MinecraftTool()

	responseDeleteCreatedPutlistBody, err := st.PutlistBodyClient.CreatePutlistBody(ctx, &docv1.CreatePutlistBodyRequest{
		PutlistNumber: putlistNumber,
		Number:        delnumber,
		ContragentId:  contragentId,
		Item:          delitem,
		TimeWith:      delparseTimeWith,
		TimeFor:       delparseTimeFor,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteCreatedPutlistBody.GetPutlistBodyId())

	responseDeletePutlistBody, err := st.PutlistBodyClient.DeletePutlistBody(ctx, &docv1.DeletePutlistBodyRequest{
		PutlistBodyId: responseDeleteCreatedPutlistBody.GetPutlistBodyId(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeletePutlistBody.GetMessage())
	require.Equal(t, responseDeletePutlistBody.GetMessage(), "deleted")
}

func Test_CRUD_PutlistBody_FailCases(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	timeWith := "2024-11-13 15:00:00"
	timeFor := "2024-11-13 17:00:00"
	number := int64(rand.IntN(100))
	item := gofakeit.MinecraftTool()

	createPutlistBodyTests := []struct {
		putlistNumber int64
		number        int64
		contragentId  int64
		item          string
		timeWith      string
		timeFor       string
		expectedErr   string
	}{
		{
			putlistNumber: emptyPutlistNumber,
			number:        number,
			contragentId:  bankAccId,
			item:          item,
			timeWith:      timeWith,
			timeFor:       timeFor,
			expectedErr:   "putlist number is required",
		},
		{
			putlistNumber: putlistNumber,
			contragentId:  contragentId,
			item:          item,
			timeWith:      timeWith,
			timeFor:       timeFor,
			expectedErr:   "number is required",
		},
		{
			putlistNumber: putlistNumber,
			number:        number,
			contragentId:  emptyContragentId,
			item:          item,
			timeWith:      timeWith,
			timeFor:       timeFor,
			expectedErr:   "contragent ID is required",
		},
		{
			putlistNumber: putlistNumber,
			number:        number,
			contragentId:  contragentId,
			item:          "",
			timeWith:      timeWith,
			timeFor:       timeFor,
			expectedErr:   "item is required",
		},
		{
			putlistNumber: putlistNumber,
			number:        number,
			contragentId:  contragentId,
			item:          item,
			timeWith:      "",
			timeFor:       timeFor,
			expectedErr:   "time with is required",
		},
		{
			putlistNumber: putlistNumber,
			number:        number,
			contragentId:  contragentId,
			item:          item,
			timeWith:      timeWith,
			timeFor:       "",
			expectedErr:   "time for is required",
		},
		{
			putlistNumber: putlistNumber,
			number:        number,
			contragentId:  contragentId,
			item:          item,
			timeWith:      timeWith,
			timeFor:       "30-30-2034",
			expectedErr:   "invalid datetime format",
		},
		{
			putlistNumber: putlistNumber,
			number:        number,
			contragentId:  contragentId,
			item:          item,
			timeWith:      timeWith,
			timeFor:       "dsfdf",
			expectedErr:   "invalid datetime format",
		},
	}

	for _, test := range createPutlistBodyTests {
		t.Run("CreatePutlistBodyFailCases", func(t *testing.T) {
			responseCreatePutlistBody, err := st.PutlistBodyClient.CreatePutlistBody(ctx, &docv1.CreatePutlistBodyRequest{
				PutlistNumber: test.putlistNumber,
				Number:        test.number,
				ContragentId:  test.contragentId,
				Item:          test.item,
				TimeWith:      test.timeWith,
				TimeFor:       test.timeFor,
			})
			require.Error(t, err)
			require.Empty(t, responseCreatePutlistBody.GetPutlistBodyId())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	responseCreatePutlistBody, err := st.PutlistBodyClient.CreatePutlistBody(ctx, &docv1.CreatePutlistBodyRequest{
		PutlistNumber: putlistNumber,
		Number:        number,
		ContragentId:  contragentId,
		Item:          item,
		TimeWith:      timeWith,
		TimeFor:       timeFor,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreatePutlistBody.GetPutlistBodyId())

	respPutBodyId := responseCreatePutlistBody.GetPutlistBodyId()
	updateNumber := int64(rand.IntN(100))
	updateContragentId := contragentId
	updateItem := gofakeit.MinecraftTool()
	updateTimeWith := timeWith
	updateTimeFor := timeFor
	notFoundPutBodyId := int64(gofakeit.IntRange(1000, 2000))
	emptyContrId := emptyContragentId
	emptyPutBodyId := int64(0)
	emptyNumber := int64(0)
	emptyItem := ""
	emptyTimeWith := ""
	emptyTimeFor := ""
	invTimeWith := "invalid"
	invTimeFor := "invalid"
	outRangeTimeWith := "40-40-2023"
	outRangeTimeFor := "12-40-2024"

	t.Run("UpdatePutlistBodyFailCases №1", func(t *testing.T) {
		responseUpdatePutlistBody, err := st.PutlistBodyClient.UpdatePutlistBody(ctx, &docv1.UpdatePutlistBodyRequest{
			Number:       &updateNumber,
			ContragentId: &updateContragentId,
			Item:         &updateItem,
			TimeWith:     &updateTimeWith,
			TimeFor:      &updateTimeFor,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlistBody.GetMessage())
		require.ErrorContains(t, err, "putlist body ID is required")
	})

	t.Run("UpdatePutlistBodyFailCases №2", func(t *testing.T) {
		responseUpdatePutlistBody, err := st.PutlistBodyClient.UpdatePutlistBody(ctx, &docv1.UpdatePutlistBodyRequest{
			PutlistBodyId: emptyPutBodyId,
			Number:        &updateNumber,
			ContragentId:  &updateContragentId,
			Item:          &updateItem,
			TimeWith:      &updateTimeWith,
			TimeFor:       &updateTimeFor,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlistBody.GetMessage())
		require.ErrorContains(t, err, "putlist body ID is required")
	})

	t.Run("UpdatePutlistBodyFailCases №3", func(t *testing.T) {
		responseUpdatePutlistBody, err := st.PutlistBodyClient.UpdatePutlistBody(ctx, &docv1.UpdatePutlistBodyRequest{
			PutlistBodyId: notFoundPutBodyId,
			Number:        &updateNumber,
			ContragentId:  &updateContragentId,
			Item:          &updateItem,
			TimeWith:      &updateTimeWith,
			TimeFor:       &updateTimeFor,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlistBody.GetMessage())
		require.ErrorContains(t, err, "putlist body not found")
	})

	t.Run("UpdatePutlistBodyFailCases №4", func(t *testing.T) {
		responseUpdatePutlistBody, err := st.PutlistBodyClient.UpdatePutlistBody(ctx, &docv1.UpdatePutlistBodyRequest{
			PutlistBodyId: respPutBodyId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlistBody.GetMessage())
		require.ErrorContains(t, err, "updating structure has no values")
	})

	t.Run("UpdatePutlistBodyFailCases №5", func(t *testing.T) {
		responseUpdatePutlistBody, err := st.PutlistBodyClient.UpdatePutlistBody(ctx, &docv1.UpdatePutlistBodyRequest{
			PutlistBodyId: respPutBodyId,
			Number:        &emptyNumber,
			ContragentId:  &emptyContrId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlistBody.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdatePutlistBodyFailCases №6", func(t *testing.T) {
		responseUpdatePutlistBody, err := st.PutlistBodyClient.UpdatePutlistBody(ctx, &docv1.UpdatePutlistBodyRequest{
			PutlistBodyId: respPutBodyId,
			Item:          &emptyItem,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlistBody.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdatePutlistBodyFailCases №7", func(t *testing.T) {
		responseUpdatePutlistBody, err := st.PutlistBodyClient.UpdatePutlistBody(ctx, &docv1.UpdatePutlistBodyRequest{
			PutlistBodyId: respPutBodyId,
			TimeWith:      &emptyTimeWith,
			TimeFor:       &emptyTimeFor,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlistBody.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdatePutlistBodyFailCases №8", func(t *testing.T) {
		responseUpdatePutlistBody, err := st.PutlistBodyClient.UpdatePutlistBody(ctx, &docv1.UpdatePutlistBodyRequest{
			PutlistBodyId: respPutBodyId,
			TimeWith:      &invTimeWith,
			TimeFor:       &invTimeFor,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlistBody.GetMessage())
		require.ErrorContains(t, err, "invalid datetime format")
	})

	t.Run("UpdatePutlistBodyFailCases №9", func(t *testing.T) {
		responseUpdatePutlistBody, err := st.PutlistBodyClient.UpdatePutlistBody(ctx, &docv1.UpdatePutlistBodyRequest{
			PutlistBodyId: respPutBodyId,
			TimeWith:      &outRangeTimeWith,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlistBody.GetMessage())
		require.ErrorContains(t, err, "invalid datetime format")
	})

	t.Run("UpdatePutlistBodyFailCases №10", func(t *testing.T) {
		responseUpdatePutlistBody, err := st.PutlistBodyClient.UpdatePutlistBody(ctx, &docv1.UpdatePutlistBodyRequest{
			PutlistBodyId: respPutBodyId,
			TimeFor:       &outRangeTimeFor,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlistBody.GetMessage())
		require.ErrorContains(t, err, "invalid datetime format")
	})

	deletePutlistBodyTests := []struct {
		putlistBodyId int64
		expectedErr   string
	}{
		{
			putlistBodyId: 0,
			expectedErr:   "putlist body ID is required",
		},
		{
			putlistBodyId: int64(gofakeit.IntRange(1000, 2000)),
			expectedErr:   "putlist body not found",
		},
	}

	for _, test := range deletePutlistBodyTests {
		t.Run("DeletePutlistBodyFailCases", func(t *testing.T) {
			responseDeletePutlistBody, err := st.PutlistBodyClient.DeletePutlistBody(ctx, &docv1.DeletePutlistBodyRequest{
				PutlistBodyId: test.putlistBodyId,
			})
			require.Error(t, err)
			require.Empty(t, responseDeletePutlistBody.GetMessage())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	getPutlistBodiesTests := []struct {
		putlistNumber int64
		expectedErr   string
	}{
		{
			putlistNumber: 0,
			expectedErr:   "putlist number is required",
		},
	}

	for _, test := range getPutlistBodiesTests {
		t.Run("GetPutlistBodiesFailCases", func(t *testing.T) {
			responseGetPutlistBodies, err := st.PutlistBodyClient.GetPutlistBodies(ctx, &docv1.GetPutlistBodiesRequest{
				PutlistNumber: test.putlistNumber,
			})
			require.Error(t, err)
			require.Empty(t, responseGetPutlistBodies.GetPutlistBodies())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

}
