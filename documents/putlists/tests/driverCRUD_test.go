package tests

import (
	"testing"

	docv1 "github.com/DavidG9999/DMS/api/grpc/document_api/gen/go"
	"github.com/DavidG9999/DMS/documents/putlists/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func Test_CRUD_Driver_HappyPath(t *testing.T) {

	ctx, st := suite.NewSuite(t)

	fullname := gofakeit.Name()
	license := randomLicense()
	class := licenseClass()

	responseCreateDriver, err := st.DriverClient.CreateDriver(ctx, &docv1.CreateDriverRequest{
		FullName: fullname,
		License:  license,
		Class:    class,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateDriver.GetDriverId())

	responseGetDrivers, err := st.DriverClient.GetDrivers(ctx, &docv1.GetDriversRequest{})
	require.NoError(t, err)
	require.NotEmpty(t, responseGetDrivers.GetDrivers())

	driverfullname := gofakeit.Name()
	driverlicense := randomLicense()
	driverclass := licenseClass()
	t.Run("UpdateDriverTest №1", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: responseCreateDriver.GetDriverId(),
			FullName: &driverfullname,
			License:  &driverlicense,
			Class:    &driverclass,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateDriver.GetMessage())
		require.Equal(t, responseUpdateDriver.GetMessage(), "updated driver")
	})

	t.Run("UpdateDriverTest №2", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: responseCreateDriver.GetDriverId(),
			FullName: &driverfullname,
			License:  &driverlicense,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateDriver.GetMessage())
		require.Equal(t, responseUpdateDriver.GetMessage(), "updated driver")
	})

	t.Run("UpdateDriverTest №3", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: responseCreateDriver.GetDriverId(),
			FullName: &driverfullname,
			Class:    &driverclass,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateDriver.GetMessage())
		require.Equal(t, responseUpdateDriver.GetMessage(), "updated driver")
	})

	t.Run("UpdateDriverTest №4", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: responseCreateDriver.GetDriverId(),
			License:  &driverlicense,
			Class:    &driverclass,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateDriver.GetMessage())
		require.Equal(t, responseUpdateDriver.GetMessage(), "updated driver")
	})

	t.Run("UpdateDriverTest №5", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: responseCreateDriver.GetDriverId(),
			FullName: &driverfullname,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateDriver.GetMessage())
		require.Equal(t, responseUpdateDriver.GetMessage(), "updated driver")
	})

	t.Run("UpdateDriverTest №6", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: responseCreateDriver.GetDriverId(),
			License:  &driverlicense,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateDriver.GetMessage())
		require.Equal(t, responseUpdateDriver.GetMessage(), "updated driver")
	})

	t.Run("UpdateDriverTest №7", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: responseCreateDriver.GetDriverId(),
			Class:    &driverclass,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseUpdateDriver.GetMessage())
		require.Equal(t, responseUpdateDriver.GetMessage(), "updated driver")
	})

	delFullName := gofakeit.Name()
	delLicense := randomLicense()
	delClass := licenseClass()

	responseDeleteCreatedDriver, err := st.DriverClient.CreateDriver(ctx, &docv1.CreateDriverRequest{
		FullName: delFullName,
		License:  delLicense,
		Class:    delClass,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteCreatedDriver.GetDriverId())

	responseDeleteDriver, err := st.DriverClient.DeleteDriver(ctx, &docv1.DeleteDriverRequest{
		DriverId: responseDeleteCreatedDriver.GetDriverId(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteDriver.GetMessage())
	require.Equal(t, responseDeleteDriver.GetMessage(), "deleted")
}

func Test_CRUD_Driver_FailCases(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	fullname := gofakeit.Company()
	license := randomLicense()
	class := licenseClass()

	t.Run("CreateDriverDublicateEmailTest", func(t *testing.T) {
		responseCreateDriver, err := st.DriverClient.CreateDriver(ctx, &docv1.CreateDriverRequest{
			FullName: fullname,
			License:  license,
			Class:    class,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseCreateDriver.GetDriverId())

		responseDubleCreateDriver, err := st.DriverClient.CreateDriver(ctx, &docv1.CreateDriverRequest{
			FullName: fullname,
			License:  license,
			Class:    class,
		})
		require.Error(t, err)
		require.Empty(t, responseDubleCreateDriver.GetDriverId())
		require.ErrorContains(t, err, "driver already exists")
	})

	createDriverTests := []struct {
		fullname    string
		license     string
		class       string
		expectedErr string
	}{
		{
			fullname:    "",
			license:     randomLicense(),
			class:       licenseClass(),
			expectedErr: "fullname is required",
		},
		{
			license:     randomLicense(),
			class:       licenseClass(),
			expectedErr: "fullname is required",
		},
		{
			fullname:    gofakeit.Name(),
			license:     "",
			class:       licenseClass(),
			expectedErr: "license is required",
		},
		{
			fullname:    gofakeit.Name(),
			class:       licenseClass(),
			expectedErr: "license is required",
		},
		{
			fullname:    gofakeit.Name(),
			license:     randomLicense(),
			class:       "",
			expectedErr: "class is required",
		},
		{
			fullname:    gofakeit.Name(),
			license:     randomLicense(),
			expectedErr: "class is required",
		},
		{
			fullname:    gofakeit.Name(),
			license:     invalidLicenseMin(),
			class:       licenseClass(),
			expectedErr: "invalid field format: license",
		},
		{
			fullname:    gofakeit.Name(),
			license:     invalidLicenseMax(),
			class:       licenseClass(),
			expectedErr: "invalid field format: license",
		},
	}

	for _, test := range createDriverTests {
		t.Run("CreateDriverFailCases", func(t *testing.T) {
			responseCreateDriver, err := st.DriverClient.CreateDriver(ctx, &docv1.CreateDriverRequest{
				FullName: test.fullname,
				License:  test.license,
				Class:    test.class,
			})
			require.Error(t, err)
			require.Empty(t, responseCreateDriver.GetDriverId())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	beforeUpFullName := gofakeit.Name()
	beforeUpLicense := randomLicense()
	beforeUpClass := licenseClass()

	responseCreateDriver, err := st.DriverClient.CreateDriver(ctx, &docv1.CreateDriverRequest{
		FullName: beforeUpFullName,
		License:  beforeUpLicense,
		Class:    beforeUpClass,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateDriver.GetDriverId())

	respDriverId := responseCreateDriver.GetDriverId()
	updateName := gofakeit.Name()
	updateLicense := randomLicense()
	updateClass := licenseClass()
	emptyDriverId := int64(0)
	emptyName := ""
	emptyLicense := ""
	emptyClass := ""
	notFoundDriverId := int64(gofakeit.IntRange(1000, 2000))
	invLicenseMin := invalidLicenseMin()
	invLicenseMax := invalidLicenseMax()

	t.Run("UpdateDriverFailCases №1", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			FullName: &updateName,
			License:  &updateLicense,
			Class:    &updateClass,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDriver.GetMessage())
		require.ErrorContains(t, err, "driver ID is required")
	})

	t.Run("UpdateDriverFailCases №2", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: emptyDriverId,
			FullName: &updateName,
			License:  &updateLicense,
			Class:    &updateClass,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDriver.GetMessage())
		require.ErrorContains(t, err, "driver ID is required")
	})

	t.Run("UpdateDriverFailCases №3", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: notFoundDriverId,
			FullName: &updateName,
			License:  &updateLicense,
			Class:    &updateClass,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDriver.GetMessage())
		require.ErrorContains(t, err, "driver not found")
	})

	t.Run("UpdateDriverFailCases №4", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: respDriverId,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDriver.GetMessage())
		require.ErrorContains(t, err, "updating structure has no values")
	})

	t.Run("UpdateDriverFailCases №5", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: respDriverId,
			FullName: &emptyName,
			License:  &emptyLicense,
			Class:    &emptyClass,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDriver.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateDriverFailCases №6", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: respDriverId,
			FullName: &emptyName,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDriver.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateDriverFailCases №7", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: respDriverId,
			License:  &emptyLicense,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDriver.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateDriverFailCases №8", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: respDriverId,
			Class:    &emptyClass,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDriver.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateDriverFailCases №9", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: respDriverId,
			License:  &emptyLicense,
			Class:    &emptyClass,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDriver.GetMessage())
		require.ErrorContains(t, err, "updating structure has empty values")
	})

	t.Run("UpdateDriverFailCases №10", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: respDriverId,
			FullName: &emptyName,
			License:  &invLicenseMin,
			Class:    &updateClass,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDriver.GetMessage())
		require.ErrorContains(t, err, "invalid field format: license")
	})

	t.Run("UpdateDriverFailCases №11", func(t *testing.T) {
		responseUpdateDriver, err := st.DriverClient.UpdateDriver(ctx, &docv1.UpdateDriverRequest{
			DriverId: respDriverId,
			License:  &invLicenseMax,
		})
		require.Error(t, err)
		require.Empty(t, responseUpdateDriver.GetMessage())
		require.ErrorContains(t, err, "invalid field format: license")
	})

	deleteDriverTests := []struct {
		driverID    int64
		expectedErr string
	}{
		{
			driverID:    0,
			expectedErr: "driver ID is required",
		},
		{
			expectedErr: "driver ID is required",
		},
		{
			driverID:    int64(gofakeit.IntRange(1000, 2000)),
			expectedErr: "driver not found",
		},
	}

	for _, test := range deleteDriverTests {
		t.Run("DeleteDriverFailCases", func(t *testing.T) {
			responseDeleteDriver, err := st.DriverClient.DeleteDriver(ctx, &docv1.DeleteDriverRequest{
				DriverId: test.driverID,
			})
			require.Error(t, err)
			require.Empty(t, responseDeleteDriver.GetMessage())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

}
