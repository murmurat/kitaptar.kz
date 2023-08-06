package jwttokentest

import (
	"github.com/stretchr/testify/require"
	"one-lab/pkg/jwttoken"
	"testing"
	"time"
)

func TestJWTToken(t *testing.T) {
	jwt := jwttoken.New("test_key")
	email := "mail@gmail.com"

	token, err := jwt.CreateToken(email, time.Minute*15)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEqual(t, token, "")
	require.NotNil(t, token)
}
