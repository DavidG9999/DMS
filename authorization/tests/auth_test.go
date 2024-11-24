package tests

import (
	"testing"
	"time"

	authgrpcv1 "github.com/DavidG9999/DMS/api/grpc/auth_api/gen/go"
	authjwt "github.com/DavidG9999/DMS/authorization/internal/lib/jwt"
	"github.com/DavidG9999/DMS/authorization/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

const (
	passDefaultLen = 10
	secretKey      = "DMS_Microservices_system1"
)

func Test_Auth_HappyPath(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	name := gofakeit.Username()
	email := gofakeit.Email()
	password := randomPassword()

	responseSignUp, err := st.AuthClient.SignUp(ctx, &authgrpcv1.SignUpRequest{
		Name:     name,
		Email:    email,
		Password: password,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseSignUp.GetUserId())

	responseSignIn, err := st.AuthClient.SignIn(ctx, &authgrpcv1.SignInRequest{
		Email:    email,
		Password: password,
	})
	require.NoError(t, err)
	require.NotEmpty(t, responseSignIn.GetToken())

	token := responseSignIn.GetToken()
	loginTime := time.Now()
	tokenParsed, err := jwt.ParseWithClaims(token, &authjwt.TokenClaims{} ,func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	require.NoError(t, err)

	claims, ok := tokenParsed.Claims.(*authjwt.TokenClaims)
	require.True(t, ok)

	require.Equal(t, responseSignUp.GetUserId(), claims.UserID)
	const deltaSeconds = 1

	require.InDelta(t, loginTime.Unix(), claims.IssuedAt, deltaSeconds)
	require.InDelta(t, loginTime.Add(st.Cfg.TokenTTL).Unix(), claims.ExpiresAt, deltaSeconds)
}

func Test_Auth_FailCases(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	name := gofakeit.Username()
	email := gofakeit.Email()
	password := randomPassword()

	t.Run("Test_Dublicate_Credentials", func(t *testing.T) {
		responseSignUp, err := st.AuthClient.SignUp(ctx, &authgrpcv1.SignUpRequest{
			Name:     name,
			Email:    email,
			Password: password,
		})
		require.NoError(t, err)
		require.NotEmpty(t, responseSignUp.GetUserId())

		responseDublSignUp, err := st.AuthClient.SignUp(ctx, &authgrpcv1.SignUpRequest{
			Name:     name,
			Email:    email,
			Password: password,
		})
		require.Error(t, err)
		require.Empty(t, responseDublSignUp.GetUserId())
		require.ErrorContains(t, err, "user already exists")

	})

	signUpTests := []struct {
		Name        string
		Email       string
		Password    string
		ExpectedErr string
	}{
		{
			Name:        "",
			Email:       gofakeit.Email(),
			Password:    randomPassword(),
			ExpectedErr: "name is required",
		},
		{
			Name:        gofakeit.Name(),
			Email:       "",
			Password:    randomPassword(),
			ExpectedErr: "email is required",
		},
		{
			Name:        gofakeit.Username(),
			Email:       gofakeit.Email(),
			Password:    "",
			ExpectedErr: "password is required",
		},
	}

	for _, test := range signUpTests {
		t.Run("SignUp_FailCases", func(t *testing.T) {
			responseSignUp, err := st.AuthClient.SignUp(ctx, &authgrpcv1.SignUpRequest{
				Name:     test.Name,
				Email:    test.Email,
				Password: test.Password,
			})
			require.Error(t, err)
			require.Empty(t, responseSignUp.GetUserId())
			require.ErrorContains(t, err, test.ExpectedErr)
		})
	}

	signInTests := []struct {
		Email       string
		Password    string
		ExpectedErr string
	}{
		{
			Email:       "",
			Password:    randomPassword(),
			ExpectedErr: "email is required",
		}, {
			Email:       gofakeit.Email(),
			Password:    "",
			ExpectedErr: "password is required",
		},
	}

	for _, test := range signInTests {
		t.Run("SignIn_FailCases", func(t *testing.T) {
			responseSignIn, err := st.AuthClient.SignIn(ctx, &authgrpcv1.SignInRequest{
				Email:    test.Email,
				Password: test.Password,
			})
			require.Error(t, err)
			require.Empty(t, responseSignIn.GetToken())
			require.Error(t, err, test.ExpectedErr)
		})
	}

}

func randomPassword() string {
	password := gofakeit.Password(true, true, true, true, true, passDefaultLen)
	return password
}
