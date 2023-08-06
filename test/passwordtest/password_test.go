package passwordtest

import (
	"github.com/stretchr/testify/require"
	"one-lab/pkg/util"
	"testing"
)

const password = "testpassword"

func TestHashPassword(t *testing.T) {
	hashedpassword, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedpassword)
	require.NotEqual(t, hashedpassword, "")
	require.NotNil(t, hashedpassword)

}

func TestCheckPassword(t *testing.T) {
	hashedpassword, err := util.HashPassword(password)
	err = util.CheckPassword(password, hashedpassword)
	require.NoError(t, err)
	require.Nil(t, err)

	err = util.CheckPassword("incorrect_password", hashedpassword)
	require.Error(t, err)
	require.NotNil(t, err)
}
