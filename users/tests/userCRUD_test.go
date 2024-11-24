package tests

import (
	"testing"

	userv1 "github.com/DavidG9999/DMS/api/grpc/user_api/gen/go"
	"github.com/DavidG9999/DMS/users/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

const (
	passDefaultLen = 10
)

func Test_CRUD_User_HappyPath(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	name := gofakeit.Username()
	email := gofakeit.Email()
	passwordHash := randomPassword()

	responseCreateUser, err := st.UserClient.CreateUser(ctx, &userv1.CreateUserRequest{
		Name:         name,
		Email:        email,
		PasswordHash: string(passwordHash),
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseCreateUser.GetUserId())

	responseGetUser, err := st.UserClient.GetUser(ctx, &userv1.GetUserRequest{
		Email: email,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseGetUser.GetUserId())
	require.NotEmpty(t, responseGetUser.GetName())
	require.NotEmpty(t, responseGetUser.GetEmail())
	require.NotEmpty(t, responseGetUser.GetPasswordHash())

	require.Equal(t, responseGetUser.GetUserId(), responseCreateUser.GetUserId())
	require.Equal(t, responseGetUser.GetName(), name)
	require.Equal(t, responseGetUser.GetEmail(), email)
	require.Equal(t, responseGetUser.GetPasswordHash(), string(passwordHash))

	responseGetUserById, err := st.UserClient.GetUserById(ctx, &userv1.GetUserByIdRequest{
		UserId: responseCreateUser.GetUserId(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseGetUserById.GetName())
	require.NotEmpty(t, responseGetUserById.GetEmail())

	require.Equal(t, responseGetUserById.GetName(), name)
	require.Equal(t, responseGetUserById.GetEmail(), email)

	updateName := gofakeit.Username()

	responseUpdateName, err := st.UserClient.UpdateName(ctx, &userv1.UpdateNameRequest{
		UserId:     responseCreateUser.GetUserId(),
		UpdateName: updateName,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseUpdateName.GetMessage())
	require.Equal(t, responseUpdateName.GetMessage(), "updated name")

	updateNPasswordHash := randomPassword()

	responseUpdatePasswordHash, err := st.UserClient.UpdatePassword(ctx, &userv1.UpdatePasswordRequest{
		UserId:         responseCreateUser.GetUserId(),
		UpdatePassword: string(updateNPasswordHash),
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseUpdatePasswordHash.GetMessage())
	require.Equal(t, responseUpdatePasswordHash.GetMessage(), "updated password")

	delName := gofakeit.Username()
	delEmail := gofakeit.Email()
	delPasswordHash := randomPassword()

	responseDeleteCreatedUser, err := st.UserClient.CreateUser(ctx, &userv1.CreateUserRequest{
		Name:         delName,
		Email:        delEmail,
		PasswordHash: string(delPasswordHash),
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteCreatedUser.GetUserId())

	responseDeleteUser, err := st.UserClient.DeleteUser(ctx, &userv1.DeleteUserRequest{
		UserId: responseDeleteCreatedUser.GetUserId(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseDeleteUser.GetMessage())
	require.Equal(t, responseDeleteUser.GetMessage(), "deleted")
}

func Test_CRUD_User_FailCases(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	name := gofakeit.Username()
	email := gofakeit.Email()
	passwordHash := randomPassword()

	t.Run("CreateUserDublicateEmailTest", func(t *testing.T) {
		responseCreateUser, err := st.UserClient.CreateUser(ctx, &userv1.CreateUserRequest{
			Name:         name,
			Email:        email,
			PasswordHash: string(passwordHash),
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseCreateUser.GetUserId())

		responseDubleCreateUser, err := st.UserClient.CreateUser(ctx, &userv1.CreateUserRequest{
			Name:         name,
			Email:        email,
			PasswordHash: string(passwordHash),
		})
		require.Error(t, err)
		require.Empty(t, responseDubleCreateUser.GetUserId())
		require.ErrorContains(t, err, "user already exists")
	})

	createUserTests := []struct {
		name         string
		email        string
		passwordHash string
		expectedErr  string
	}{
		{
			name:         "",
			email:        gofakeit.Email(),
			passwordHash: string(randomPassword()),
			expectedErr:  "name is required",
		},
		{
			name:         gofakeit.Username(),
			email:        "",
			passwordHash: string(randomPassword()),
			expectedErr:  "email is required",
		},
		{
			name:         gofakeit.Username(),
			email:        gofakeit.Email(),
			passwordHash: "",
			expectedErr:  "password is required",
		},
	}

	for _, test := range createUserTests {
		t.Run("CreateUserFailCases", func(t *testing.T) {
			responseCreateUser, err := st.UserClient.CreateUser(ctx, &userv1.CreateUserRequest{
				Name:         test.name,
				Email:        test.email,
				PasswordHash: test.passwordHash,
			})
			require.Error(t, err)
			require.Empty(t, responseCreateUser.GetUserId())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	getUserTests := []struct {
		email       string
		expectedErr string
	}{
		{
			email:       "",
			expectedErr: "email is required",
		},
		{
			email:       gofakeit.Email(),
			expectedErr: "user not found",
		},
	}

	for _, test := range getUserTests {
		t.Run("GetUserFailCases", func(t *testing.T) {
			responseGetUser, err := st.UserClient.GetUser(ctx, &userv1.GetUserRequest{
				Email: test.email,
			})
			require.Error(t, err)
			require.Empty(t, responseGetUser.GetUserId())
			require.Empty(t, responseGetUser.GetName())
			require.Empty(t, responseGetUser.GetEmail())
			require.Empty(t, responseGetUser.GetPasswordHash())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	getUserByIdTests := []struct {
		userID      int64
		expectedErr string
	}{
		{
			userID:      0,
			expectedErr: "user ID is required",
		},
		{
			userID:      int64(gofakeit.IntRange(100, 200)),
			expectedErr: "user not found",
		},
	}

	for _, test := range getUserByIdTests {
		t.Run("GetUserByIdFailCases", func(t *testing.T) {
			responseGetUserById, err := st.UserClient.GetUserById(ctx, &userv1.GetUserByIdRequest{
				UserId: test.userID,
			})
			require.Error(t, err)
			require.Empty(t, responseGetUserById.GetName())
			require.Empty(t, responseGetUserById.GetEmail())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	updateNameTests := []struct {
		userID      int64
		updateName  string
		expectedErr string
	}{
		{
			userID:      0,
			updateName:  string(randomPassword()),
			expectedErr: "user ID is required",
		},
		{
			userID:      int64(gofakeit.IntRange(100,200)),
			updateName:  string(randomPassword()),
			expectedErr: "user not found",
		},
	}

	for _, test := range updateNameTests {
		t.Run("UpdateNameFailCases", func(t *testing.T) {
			responseUpdateName, err := st.UserClient.UpdateName(ctx, &userv1.UpdateNameRequest{
				UserId:     test.userID,
				UpdateName: test.updateName,
			})
			require.Error(t, err)
			require.Empty(t, responseUpdateName.GetMessage())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	updatePasswordTests := []struct {
		userID         int64
		UpdatePassword string
		expectedErr    string
	}{
		{
			userID:         0,
			UpdatePassword: string(randomPassword()),
			expectedErr:    "user ID is required",
		},
		{
			userID:         int64(gofakeit.IntRange(200, 400)),
			UpdatePassword: string(randomPassword()),
			expectedErr:    "user not found",
		},
	}

	for _, test := range updatePasswordTests {
		t.Run("UpdatePasswordFailCases", func(t *testing.T) {
			responseUpdatePassword, err := st.UserClient.UpdatePassword(ctx, &userv1.UpdatePasswordRequest{
				UserId:         test.userID,
				UpdatePassword: test.UpdatePassword,
			})
			require.Error(t, err)
			require.Empty(t, responseUpdatePassword.GetMessage())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

	deleteUserTests := []struct {
		userID      int64
		expectedErr string
	}{
		{
			userID:      0,
			expectedErr: "user ID is required",
		},
		{
			userID:      int64(gofakeit.IntRange(200, 400)),
			expectedErr: "user not found",
		},
	}

	for _, test := range deleteUserTests {
		t.Run("DeleteUserFailCases", func(t *testing.T) {
			responseDeleteUser, err := st.UserClient.DeleteUser(ctx, &userv1.DeleteUserRequest{
				UserId: test.userID,
			})
			require.Error(t, err)
			require.Empty(t, responseDeleteUser.GetMessage())
			require.ErrorContains(t, err, test.expectedErr)
		})
	}

}

func randomPassword() []byte {
	password := gofakeit.Password(true, true, true, true, true, passDefaultLen)

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return passwordHash
}
