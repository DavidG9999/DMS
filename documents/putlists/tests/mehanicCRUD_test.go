package tests

import (
	"testing"

	docv1 "github.com/DavidG9999/DMS/api/grpc/document_api/gen/go"
	"github.com/DavidG9999/DMS/documents/putlists/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func Test_CRUD_Mehanic_HappyPath(t *testing.T) {

	ctx, st := suite.NewSuite(t)

	fullname := gofakeit.Name()

	responseCreateMehanic, err := st.MehanicClient.CreateMehanic(ctx, &docv1.CreateMehanicRequest{
		FullName: fullname,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateMehanic.GetMehanicId())

	responseGetMehanics, err := st.MehanicClient.GetMehanics(ctx, &docv1.GetMehanicsRequest{})
	require.NoError(t, err)
	require.NotEmpty(t, responseGetMehanics.GetMehanics())

	mehanicFullName := gofakeit.Name()
	t.Run("UpdateMehanicTest", func(t *testing.T) {
		responseUpdateMehanic, err := st.MehanicClient.UpdateMehanic(ctx, &docv1.UpdateMehanicRequest{
			MehanicId: responseCreateMehanic.GetMehanicId(),
			FullName:  &mehanicFullName,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateMehanic.GetMessage())
		require.Equal(t, responseUpdateMehanic.GetMessage(), "updated mehanic")
	})

	delFullName := gofakeit.Name()

	responseDeleteCreatedMehanic, err := st.MehanicClient.CreateMehanic(ctx, &docv1.CreateMehanicRequest{
		FullName: delFullName,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteCreatedMehanic.GetMehanicId())

	responseDeleteMehanic, err := st.MehanicClient.DeleteMehanic(ctx, &docv1.DeleteMehanicRequest{
		MehanicId: responseDeleteCreatedMehanic.GetMehanicId(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteMehanic.GetMessage())
	require.Equal(t, responseDeleteMehanic.GetMessage(), "deleted")
}

func Test_CRUD_Mehanic_FailCases(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	createMehanicTests := []struct {
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

	for _, test := range createMehanicTests {
		t.Run("CreateMehanicFailCases", func(t *testing.T) {
			responseCreateMehanic, err := st.MehanicClient.CreateMehanic(ctx, &docv1.CreateMehanicRequest{
				FullName: test.fullname,
			})
			require.Error(t, err)
			require.Empty(t, responseCreateMehanic.GetMehanicId())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	beforeUpFullName := gofakeit.Name()

	responseCreateMehanic, err := st.MehanicClient.CreateMehanic(ctx, &docv1.CreateMehanicRequest{
		FullName: beforeUpFullName,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateMehanic.GetMehanicId())

	respMehanicId := responseCreateMehanic.GetMehanicId()
	updateFullName := gofakeit.Name()
	emptyFullName := ""
	notFounMehanicId := int64(gofakeit.IntRange(1000, 2000))

	t.Run("UpdateMehanicFailCases №1", func(t *testing.T) {
		responseUpdateMehanic, err := st.MehanicClient.UpdateMehanic(ctx, &docv1.UpdateMehanicRequest{
			MehanicId: 0,
			FullName:  &updateFullName,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateMehanic.GetMessage())
		require.ErrorContains(t, err, "mehanic ID is required")
	})

	t.Run("UpdateMehanicFailCases №2", func(t *testing.T) {
		responseUpdateMehanic, err := st.MehanicClient.UpdateMehanic(ctx, &docv1.UpdateMehanicRequest{
			FullName: &updateFullName,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateMehanic.GetMessage())
		require.ErrorContains(t, err, "mehanic ID is required")
	})

	t.Run("UpdateMehanicFailCases №3", func(t *testing.T) {
		responseUpdateMehanic, err := st.MehanicClient.UpdateMehanic(ctx, &docv1.UpdateMehanicRequest{
			MehanicId: notFounMehanicId,
			FullName:  &updateFullName,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateMehanic.GetMessage())
		require.ErrorContains(t, err, "mehanic not found")
	})

	t.Run("UpdateMehanicFailCases №4", func(t *testing.T) {
		responseUpdateMehanic, err := st.MehanicClient.UpdateMehanic(ctx, &docv1.UpdateMehanicRequest{
			MehanicId: respMehanicId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateMehanic.GetMessage())
		require.ErrorContains(t, err, "updating structure has no values")
	})

	t.Run("UpdateMehanicFailCases №5", func(t *testing.T) {
		responseUpdateMehanic, err := st.MehanicClient.UpdateMehanic(ctx, &docv1.UpdateMehanicRequest{
			MehanicId: respMehanicId,
			FullName:  &emptyFullName,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateMehanic.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	deleteMehanicTests := []struct {
		mehanicID   int64
		expectedErr string
	}{
		{
			mehanicID:   0,
			expectedErr: "mehanic ID is required",
		},
		{
			expectedErr: "mehanic ID is required",
		},
		{
			mehanicID:   int64(gofakeit.IntRange(1000, 2000)),
			expectedErr: "mehanic not found",
		},
	}

	for _, test := range deleteMehanicTests {
		t.Run("DeleteMehanicFailCases", func(t *testing.T) {
			responseDeleteMehanic, err := st.MehanicClient.DeleteMehanic(ctx, &docv1.DeleteMehanicRequest{
				MehanicId: test.mehanicID,
			})
			require.Error(t, err)
			require.Empty(t, responseDeleteMehanic.GetMessage())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

}
