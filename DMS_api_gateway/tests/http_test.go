package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	parsejwt "github.com/DavidG9999/DMS/DMS_api_gateway/internal/lib/jwt"
	"github.com/DavidG9999/DMS/DMS_api_gateway/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

const (
	passDefaultLen = 10
	secretKey      = "DMS_Microservices_system1"
	authHeader     = "Authorization"
)

func Test_HappyPath(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	name := gofakeit.Username()
	email := gofakeit.Email()
	password := randomPassword()

	reqSignUp := fmt.Sprintf(`{"name": "%s", "email": "%s", "password": "%s"}`, name, email, password)

	reqSignUpByte := []byte(reqSignUp)
	reqSignUpReader := bytes.NewReader(reqSignUpByte)
	respSignUp, err := st.HTTPClient.Post("http://localhost:8080/auth/sign-up", "application/json", reqSignUpReader)

	require.NoError(t, err)
	require.NotEmpty(t, respSignUp)

	var respSignUpInterface map[string]int64
	respSignUpDecoder := json.NewDecoder(respSignUp.Body)
	err = respSignUpDecoder.Decode(&respSignUpInterface)
	require.NoError(t, err)

	reqSignIn := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password)

	reqSignInByte := []byte(reqSignIn)
	reqSignInReader := bytes.NewReader(reqSignInByte)
	respSignIn, err := st.HTTPClient.Post("http://localhost:8080/auth/sign-in", "application/json", reqSignInReader)

	require.NoError(t, err)
	require.NotEmpty(t, respSignIn)

	var respSignInInterface map[string]string
	respSignInDecoder := json.NewDecoder(respSignIn.Body)
	err = respSignInDecoder.Decode(&respSignInInterface)
	require.NoError(t, err)

	token := respSignInInterface["token"]
	userId := respSignUpInterface["user_id"]

	tokenParsed, err := jwt.ParseWithClaims(token, &parsejwt.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	require.NoError(t, err)
	claims, ok := tokenParsed.Claims.(*parsejwt.TokenClaims)
	require.True(t, ok)
	require.Equal(t, userId, claims.UserID)

	bearerToken := fmt.Sprintf("Bearer %s", token)

	reqGetUser, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/user", nil)
	require.NoError(t, err)

	reqGetUser.Header.Add(authHeader, bearerToken)
	respGetUser, err := st.HTTPClient.Do(reqGetUser)

	require.NoError(t, err)
	require.NotEmpty(t, respGetUser)
	require.NotEqual(t, "401 Unauthorized", respGetUser.Status)

	updateName := gofakeit.Username()
	reqUpdateNameBody := fmt.Sprintf(`{"name": "%s"}`, updateName)

	reqUpdateNameByte := []byte(reqUpdateNameBody)
	reqUpdateNameReader := bytes.NewReader(reqUpdateNameByte)

	reqUpdateName, err := http.NewRequestWithContext(ctx, "PUT", "http://localhost:8080/user/name", reqUpdateNameReader)
	require.NoError(t, err)

	reqUpdateName.Header.Add(authHeader, bearerToken)
	respUpdateName, err := st.HTTPClient.Do(reqUpdateName)

	require.NoError(t, err)
	require.NotEmpty(t, respUpdateName)
	require.NotEqual(t, "401 Unauthorized", respUpdateName.Status)

	var respUpdateNameInterface map[string]string
	respUpdateNameDecoder := json.NewDecoder(respUpdateName.Body)
	err = respUpdateNameDecoder.Decode(&respUpdateNameInterface)
	require.NoError(t, err)
	require.Equal(t, "updated name", respUpdateNameInterface["status"])

	updatePassword := randomPassword()
	reqUpdatePasswordBody := fmt.Sprintf(`{"password": "%s"}`, updatePassword)

	reqUpdatePasswordByte := []byte(reqUpdatePasswordBody)
	reqUpdatePasswordReader := bytes.NewReader(reqUpdatePasswordByte)

	reqUpdatePassword, err := http.NewRequestWithContext(ctx, "PUT", "http://localhost:8080/user/password", reqUpdatePasswordReader)
	require.NoError(t, err)

	reqUpdatePassword.Header.Add(authHeader, bearerToken)
	respUpdatePassword, err := st.HTTPClient.Do(reqUpdatePassword)

	require.NoError(t, err)
	require.NotEmpty(t, respUpdatePassword)
	require.NotEqual(t, "401 Unauthorized", respUpdatePassword.Status)

	var respUpdatePasswordInterface map[string]string
	respUpdatePasswordDecoder := json.NewDecoder(respUpdatePassword.Body)
	err = respUpdatePasswordDecoder.Decode(&respUpdatePasswordInterface)
	require.NoError(t, err)
	require.Equal(t, "updated password", respUpdatePasswordInterface["status"])

	reqDeleteUser, err := http.NewRequestWithContext(ctx, "DELETE", "http://localhost:8080/user", nil)
	require.NoError(t, err)

	reqDeleteUser.Header.Add(authHeader, bearerToken)
	respDeleteUser, err := st.HTTPClient.Do(reqDeleteUser)

	require.NoError(t, err)
	require.NotEmpty(t, respDeleteUser)
	require.NotEqual(t, "401 Unauthorized", respDeleteUser.Status)

	var respDeleteUserInterface map[string]string
	respDeleteUserDecoder := json.NewDecoder(respDeleteUser.Body)
	err = respDeleteUserDecoder.Decode(&respDeleteUserInterface)
	require.NoError(t, err)
	require.Equal(t, "deleted", respDeleteUserInterface["status"])

}

func randomPassword() string {
	password := gofakeit.Password(true, true, true, true, true, passDefaultLen)
	return password
}
