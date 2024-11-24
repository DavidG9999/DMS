package tests

import (
	"testing"

	docv1 "github.com/DavidG9999/DMS/api/grpc/document_api/gen/go"
	"github.com/DavidG9999/DMS/documents/putlists/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

const (
	userId            int64 = 1
	emptyUserId       int64 = 0
	bankAccId         int64 = 1
	emptyBankAccId    int64 = 0
	autoId            int64 = 1
	emptyAutoId       int64 = 0
	driverId          int64 = 1
	emptyDriverId     int64 = 0
	dispetcherId      int64 = 1
	emptyDispetcherId int64 = 0
	mehanicId         int64 = 1
	emptyMehanicId    int64 = 0
	num               int64 = 1
	emptyNum          int64 = 0
)

func Test_CRUD_Putlist_HappyPath(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	dateWith := "2024-11-13"
	dateFor := "2024-11-13"
	number := int64(gofakeit.IntRange(100, 200))

	responseCreatePutlist, err := st.PutlistClient.CreatePutlist(ctx, &docv1.CreatePutlistRequest{
		UserId:        userId,
		Number:        number,
		BankAccountId: bankAccId,
		DateWith:      dateWith,
		DateFor:       dateFor,
		AutoId:        autoId,
		DriverId:      driverId,
		DispetcherId:  dispetcherId,
		MehanicId:     mehanicId,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreatePutlist.GetPutlistId())

	responseGetPutlists, err := st.PutlistClient.GetPutlists(ctx, &docv1.GetPutlistsRequest{
		UserId: userId,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseGetPutlists.GetPutlists())

	responseGetPutlistByNumber, err := st.PutlistClient.GetPutlistByNumber(ctx, &docv1.GetPutlistByNumberRequest{
		UserId: userId,
		Number: number,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseGetPutlistByNumber.GetPutlist())

	updateBankAccId := bankAccId
	updateAutoId := autoId
	updateDriverId := driverId
	updateDispetcherId := dispetcherId
	updateMehanicId := mehanicId
	t.Run("UpdatePutlistTest №1", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:        userId,
			Number:        number,
			BankAccountId: &updateBankAccId,
			DateWith:      &dateWith,
			DateFor:       &dateFor,
			AutoId:        &updateAutoId,
			DriverId:      &updateDriverId,
			DispetcherId:  &updateDispetcherId,
			MehanicId:     &updateMehanicId,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdatePutlist.GetMessage())
		require.Equal(t, responseUpdatePutlist.GetMessage(), "updated putlist")
	})

	t.Run("UpdatePutlistTest №2", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:   userId,
			Number:   number,
			DateWith: &dateWith,
			DateFor:  &dateFor,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdatePutlist.GetMessage())
		require.Equal(t, responseUpdatePutlist.GetMessage(), "updated putlist")
	})

	t.Run("UpdatePutlistTest №3", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:        userId,
			Number:        number,
			BankAccountId: &updateBankAccId,
			AutoId:        &updateAutoId,
			DriverId:      &updateDriverId,
			DispetcherId:  &updateDispetcherId,
			MehanicId:     &updateMehanicId,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdatePutlist.GetMessage())
		require.Equal(t, responseUpdatePutlist.GetMessage(), "updated putlist")
	})

	delDateWith := "2024-11-13"
	delDateFor := "2024-11-13"
	delNumber := int64(gofakeit.IntRange(1000, 2000))

	responseDeleteCreatedPutlist, err := st.PutlistClient.CreatePutlist(ctx, &docv1.CreatePutlistRequest{
		UserId:        userId,
		Number:        delNumber,
		BankAccountId: bankAccId,
		DateWith:      delDateWith,
		DateFor:       delDateFor,
		AutoId:        autoId,
		DriverId:      driverId,
		DispetcherId:  dispetcherId,
		MehanicId:     mehanicId,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteCreatedPutlist.GetPutlistId())

	responseDeletePutlist, err := st.PutlistClient.DeletePutlist(ctx, &docv1.DeletePutlistRequest{
		UserId: userId,
		Number: delNumber,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeletePutlist.GetMessage())
	require.Equal(t, responseDeletePutlist.GetMessage(), "deleted")
}

func Test_CRUD_Putlist_FailCases(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	dateWith := "2024-11-13"
	dateFor := "2024-11-13"
	number := int64(gofakeit.IntRange(100, 2000))

	t.Run("CreatePutlistDublicateEmailTest", func(t *testing.T) {
		responseCreatePutlist, err := st.PutlistClient.CreatePutlist(ctx, &docv1.CreatePutlistRequest{
			UserId:        userId,
			Number:        number,
			BankAccountId: bankAccId,
			DateWith:      dateWith,
			DateFor:       dateFor,
			AutoId:        autoId,
			DriverId:      driverId,
			DispetcherId:  dispetcherId,
			MehanicId:     mehanicId,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseCreatePutlist.GetPutlistId())

		responseDubleCreatePutlist, err := st.PutlistClient.CreatePutlist(ctx, &docv1.CreatePutlistRequest{
			UserId:        userId,
			Number:        number,
			BankAccountId: bankAccId,
			DateWith:      dateWith,
			DateFor:       dateFor,
			AutoId:        autoId,
			DriverId:      driverId,
			DispetcherId:  dispetcherId,
			MehanicId:     mehanicId,
		})
		require.Error(t, err)
		require.Empty(t, responseDubleCreatePutlist.GetPutlistId())
		require.ErrorContains(t, err, "putlist already exists")
	})

	createPutlistTests := []struct {
		userId       int64
		number       int64
		bankAccId    int64
		dateWith     string
		dateFor      string
		autoId       int64
		driverId     int64
		dispetcherId int64
		mehanicId    int64
		expectedErr  string
	}{
		{
			userId:       emptyUserId,
			number:       int64(gofakeit.IntRange(10, 20)),
			bankAccId:    bankAccId,
			dateWith:     dateWith,
			dateFor:      dateFor,
			autoId:       autoId,
			driverId:     driverId,
			dispetcherId: dispetcherId,
			mehanicId:    mehanicId,
			expectedErr:  "user ID is required",
		},
		{
			userId:       userId,
			number:       0,
			bankAccId:    bankAccId,
			dateWith:     dateWith,
			dateFor:      dateFor,
			autoId:       autoId,
			driverId:     driverId,
			dispetcherId: dispetcherId,
			mehanicId:    mehanicId,
			expectedErr:  "number is required",
		},
		{
			userId:       userId,
			number:       int64(gofakeit.IntRange(10, 20)),
			bankAccId:    emptyBankAccId,
			dateWith:     dateWith,
			dateFor:      dateFor,
			autoId:       autoId,
			driverId:     driverId,
			dispetcherId: dispetcherId,
			mehanicId:    mehanicId,
			expectedErr:  "bank account ID is required",
		},
		{
			userId:       userId,
			number:       int64(gofakeit.IntRange(10, 20)),
			bankAccId:    bankAccId,
			dateWith:     "",
			dateFor:      dateFor,
			autoId:       autoId,
			driverId:     driverId,
			dispetcherId: dispetcherId,
			mehanicId:    mehanicId,
			expectedErr:  "date with is required",
		},
		{
			userId:       userId,
			number:       int64(gofakeit.IntRange(10, 20)),
			bankAccId:    bankAccId,
			dateWith:     dateWith,
			dateFor:      "",
			autoId:       autoId,
			driverId:     driverId,
			dispetcherId: dispetcherId,
			mehanicId:    mehanicId,
			expectedErr:  "date for is required",
		},
		{
			userId:       userId,
			number:       int64(gofakeit.IntRange(10, 20)),
			bankAccId:    bankAccId,
			dateWith:     dateWith,
			dateFor:      dateFor,
			autoId:       emptyAutoId,
			driverId:     driverId,
			dispetcherId: dispetcherId,
			mehanicId:    mehanicId,
			expectedErr:  "auto ID is required",
		},
		{
			userId:       userId,
			number:       int64(gofakeit.IntRange(10, 20)),
			bankAccId:    bankAccId,
			dateWith:     dateWith,
			dateFor:      dateFor,
			autoId:       autoId,
			driverId:     emptyDriverId,
			dispetcherId: dispetcherId,
			mehanicId:    mehanicId,
			expectedErr:  "driver ID is required",
		},
		{
			userId:       userId,
			number:       int64(gofakeit.IntRange(10, 20)),
			bankAccId:    bankAccId,
			dateWith:     dateWith,
			dateFor:      dateFor,
			autoId:       autoId,
			driverId:     driverId,
			dispetcherId: emptyDispetcherId,
			mehanicId:    mehanicId,
			expectedErr:  "dispetcher ID is required",
		},
		{
			userId:       userId,
			number:       int64(gofakeit.IntRange(10, 20)),
			bankAccId:    bankAccId,
			dateWith:     dateWith,
			dateFor:      dateFor,
			autoId:       autoId,
			driverId:     driverId,
			dispetcherId: dispetcherId,
			mehanicId:    emptyMehanicId,
			expectedErr:  "mehanic ID is required",
		},
		{
			userId:       userId,
			number:       int64(gofakeit.IntRange(10, 20)),
			bankAccId:    bankAccId,
			dateWith:     "dateWith",
			dateFor:      dateFor,
			autoId:       autoId,
			driverId:     driverId,
			dispetcherId: dispetcherId,
			mehanicId:    mehanicId,
			expectedErr:  "invalid datetime format",
		},
		{
			userId:       userId,
			number:       int64(gofakeit.IntRange(10, 20)),
			bankAccId:    bankAccId,
			dateWith:     dateWith,
			dateFor:      "dateFor",
			autoId:       autoId,
			driverId:     driverId,
			dispetcherId: dispetcherId,
			mehanicId:    mehanicId,
			expectedErr:  "invalid datetime format",
		},
		{
			userId:       userId,
			number:       int64(gofakeit.IntRange(10, 20)),
			bankAccId:    bankAccId,
			dateWith:     dateWith,
			dateFor:      "30-30-2024",
			autoId:       autoId,
			driverId:     driverId,
			dispetcherId: dispetcherId,
			mehanicId:    mehanicId,
			expectedErr:  "invalid datetime format",
		},
	}

	for _, test := range createPutlistTests {
		t.Run("CreatePutlistFailCases", func(t *testing.T) {
			responseCreatePutlist, err := st.PutlistClient.CreatePutlist(ctx, &docv1.CreatePutlistRequest{
				UserId:        test.userId,
				Number:        test.number,
				BankAccountId: test.bankAccId,
				DateWith:      test.dateWith,
				DateFor:       test.dateFor,
				AutoId:        test.autoId,
				DriverId:      test.driverId,
				DispetcherId:  test.dispetcherId,
				MehanicId:     test.mehanicId,
			})
			require.Error(t, err)
			require.Empty(t, responseCreatePutlist.GetPutlistId())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	responseCreatePutlist, err := st.PutlistClient.CreatePutlist(ctx, &docv1.CreatePutlistRequest{
		UserId:        userId,
		Number:        int64(gofakeit.IntRange(10, 59)),
		BankAccountId: bankAccId,
		DateWith:      dateWith,
		DateFor:       dateFor,
		AutoId:        autoId,
		DriverId:      driverId,
		DispetcherId:  dispetcherId,
		MehanicId:     mehanicId,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreatePutlist.GetPutlistId())

	updateUserId := userId
	updateNumber := num
	updateBankAccountId := bankAccId
	updateDateWith := dateWith
	updateDateFor := dateFor
	updateAutoId := autoId
	updateDriverId := driverId
	updateDispetcherId := dispetcherId
	updateMehanicId := mehanicId
	notFoundUserId := int64(gofakeit.IntRange(1000, 2000))
	notFoundNumber := int64(gofakeit.IntRange(1000, 2000))
	emptUserId := emptyUserId
	emptNumber := emptyNum
	emptBankAccountId := emptyBankAccId
	emptAutoId := emptyAutoId
	emptDriverId := emptyDriverId
	emptDispetcherId := emptyDispetcherId
	emptMehanicId := emptyMehanicId
	emptDateWith := ""
	emptDateFor := ""
	invDateWith := "invalid"
	invDateFor := "invalid"
	outRangeDateWith := "40-40-2023"
	outRangeDateFor := "12-40-2024"

	t.Run("UpdatePutlistFailCases №1", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			Number:        updateNumber,
			BankAccountId: &updateBankAccountId,
			DateWith:      &updateDateWith,
			DateFor:       &updateDateFor,
			AutoId:        &updateAutoId,
			DriverId:      &updateDriverId,
			DispetcherId:  &updateDispetcherId,
			MehanicId:     &updateMehanicId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "user ID is required")
	})

	t.Run("UpdatePutlistFailCases №2", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:        emptUserId,
			Number:        updateNumber,
			BankAccountId: &updateBankAccountId,
			DateWith:      &updateDateWith,
			DateFor:       &updateDateFor,
			AutoId:        &updateAutoId,
			DriverId:      &updateDriverId,
			DispetcherId:  &updateDispetcherId,
			MehanicId:     &updateMehanicId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "user ID is required")
	})

	t.Run("UpdatePutlistFailCases №3", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:        updateUserId,
			Number:        emptNumber,
			BankAccountId: &updateBankAccountId,
			DateWith:      &updateDateWith,
			DateFor:       &updateDateFor,
			AutoId:        &updateAutoId,
			DriverId:      &updateDriverId,
			DispetcherId:  &updateDispetcherId,
			MehanicId:     &updateMehanicId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "number is required")
	})

	t.Run("UpdatePutlistFailCases №4", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:        updateUserId,
			BankAccountId: &updateBankAccountId,
			DateWith:      &updateDateWith,
			DateFor:       &updateDateFor,
			AutoId:        &updateAutoId,
			DriverId:      &updateDriverId,
			DispetcherId:  &updateDispetcherId,
			MehanicId:     &updateMehanicId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "number is required")
	})

	t.Run("UpdatePutlistFailCases №5", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:        notFoundUserId,
			Number:        updateNumber,
			BankAccountId: &updateBankAccountId,
			DateWith:      &updateDateWith,
			DateFor:       &updateDateFor,
			AutoId:        &updateAutoId,
			DriverId:      &updateDriverId,
			DispetcherId:  &updateDispetcherId,
			MehanicId:     &updateMehanicId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "putlist not found")
	})

	t.Run("UpdatePutlistFailCases №6", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:        notFoundUserId,
			Number:        notFoundNumber,
			BankAccountId: &updateBankAccountId,
			DateWith:      &updateDateWith,
			DateFor:       &updateDateFor,
			AutoId:        &updateAutoId,
			DriverId:      &updateDriverId,
			DispetcherId:  &updateDispetcherId,
			MehanicId:     &updateMehanicId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "putlist not found")
	})

	t.Run("UpdatePutlistFailCases №7", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId: updateUserId,
			Number: updateNumber,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "updating structure has no values")
	})

	t.Run("UpdatePutlistFailCases №8", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:   updateUserId,
			Number:   updateNumber,
			DateWith: &emptDateWith,
			DateFor:  &emptDateFor,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdatePutlistFailCases №9", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:        updateUserId,
			Number:        updateNumber,
			BankAccountId: &emptBankAccountId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdatePutlistFailCases №10", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:       updateUserId,
			Number:       updateNumber,
			AutoId:       &emptAutoId,
			DriverId:     &emptDriverId,
			DispetcherId: &emptDispetcherId,
			MehanicId:    &emptMehanicId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdatePutlistFailCases №11", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:        updateUserId,
			Number:        updateNumber,
			BankAccountId: &updateBankAccountId,
			DateWith:      &outRangeDateWith,
			DateFor:       &invDateFor,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "invalid datetime format")
	})

	t.Run("UpdatePutlistFailCases №12", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:  updateUserId,
			Number:  updateNumber,
			DateFor: &outRangeDateFor,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "invalid datetime format")
	})

	t.Run("UpdatePutlistFailCases №13", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:       updateUserId,
			Number:       updateNumber,
			DateWith:     &outRangeDateWith,
			AutoId:       &updateAutoId,
			DriverId:     &updateDriverId,
			DispetcherId: &updateDispetcherId,
			MehanicId:    &updateMehanicId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "invalid datetime format")
	})

	t.Run("UpdatePutlistFailCases №14", func(t *testing.T) {
		responseUpdatePutlist, err := st.PutlistClient.UpdatePutlist(ctx, &docv1.UpdatePutlistRequest{
			UserId:   updateUserId,
			Number:   updateNumber,
			DateWith: &invDateWith,
			DateFor:  &invDateFor,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdatePutlist.GetMessage())
		require.ErrorContains(t, err, "invalid datetime format")
	})

	deletePutlistTests := []struct {
		userId      int64
		number      int64
		expectedErr string
	}{
		{
			userId:      emptyUserId,
			number:      number,
			expectedErr: "user ID is required",
		},
		{
			userId:      userId,
			number:      0,
			expectedErr: "number is required",
		},
		{
			userId:      int64(gofakeit.IntRange(100, 200)),
			number:      int64(gofakeit.IntRange(100, 200)),
			expectedErr: "putlist not found",
		},
		{
			userId:      userId,
			number:      int64(gofakeit.IntRange(100, 200)),
			expectedErr: "putlist not found",
		},
		{
			userId:      int64(gofakeit.IntRange(100, 200)),
			number:      number,
			expectedErr: "putlist not found",
		},
	}

	for _, test := range deletePutlistTests {
		t.Run("DeletePutlistFailCases", func(t *testing.T) {
			responseDeletePutlist, err := st.PutlistClient.DeletePutlist(ctx, &docv1.DeletePutlistRequest{
				UserId: test.userId,
				Number: test.number,
			})
			require.Error(t, err)
			require.Empty(t, responseDeletePutlist.GetMessage())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	getPutlistsTests := []struct {
		userId      int64
		expectedErr string
	}{
		{
			userId:      emptyUserId,
			expectedErr: "user ID is required",
		},
	}

	for _, test := range getPutlistsTests {
		t.Run("GetPutlistsFailCases", func(t *testing.T) {
			responseGetPutlists, err := st.PutlistClient.GetPutlists(ctx, &docv1.GetPutlistsRequest{
				UserId: test.userId,
			})
			require.Error(t, err)
			require.Empty(t, responseGetPutlists.GetPutlists())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	getPutlistByNumberTests := []struct {
		userId      int64
		number      int64
		expectedErr string
	}{
		{
			userId:      emptyUserId,
			number:      number,
			expectedErr: "user ID is required",
		},
		{
			userId:      userId,
			number:      0,
			expectedErr: "number is required",
		},
		{
			userId:      int64(gofakeit.IntRange(100, 200)),
			number:      int64(gofakeit.IntRange(100, 200)),
			expectedErr: "putlist by this number not found",
		},
		{
			userId:      int64(gofakeit.IntRange(100, 200)),
			number:      number,
			expectedErr: "putlist by this number not found",
		},
		{
			userId:      userId,
			number:      int64(gofakeit.IntRange(100, 200)),
			expectedErr: "putlist by this number not found",
		},
	}

	for _, test := range getPutlistByNumberTests {
		t.Run("GetPutliststByNumberFailCases", func(t *testing.T) {
			responseGetPutlistByNumber, err := st.PutlistClient.GetPutlistByNumber(ctx, &docv1.GetPutlistByNumberRequest{
				UserId: test.userId,
				Number: test.number,
			})
			require.Error(t, err)
			require.Empty(t, responseGetPutlistByNumber.GetPutlist())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}
}
