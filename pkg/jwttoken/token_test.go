package jwttoken

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJWTToken(t *testing.T) {

	jwt := New("test_key")
	email := "mail@gmail.com"
	userID, _ := uuid.NewRandom()

	token, err := jwt.CreateToken(userID.String(), email, time.Minute*15)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEqual(t, token, "")
	require.NotNil(t, token)

	claim, err := jwt.ValidateToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, claim)
	require.Equal(t, claim.UserID, userID.String())
	require.Equal(t, claim.Email, email)
}

func TestExpired(t *testing.T) {
	jwt := New("test_key")
	email := "mail@gmail.com"
	userID, _ := uuid.NewRandom()

	token, err := jwt.CreateToken(userID.String(), email, -time.Minute*15)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	claim, err := jwt.ValidateToken(token)
	require.Error(t, err)
	require.Nil(t, claim)
}
