package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const password = "testpassword"

func TestHashPassword(t *testing.T) {
	hashedpassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedpassword)
	require.NotEqual(t, hashedpassword, "")
	require.NotNil(t, hashedpassword)

}

func TestCheckPassword(t *testing.T) {
	hashedpassword, err := HashPassword(password)
	err = CheckPassword(password, hashedpassword)
	require.NoError(t, err)
	require.Nil(t, err)

	err = CheckPassword("incorrect_password", hashedpassword)
	require.Error(t, err)
	require.NotNil(t, err)
}
