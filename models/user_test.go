package models

import(
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	password := "test-password"
	user := User{
		ID: 0,
		FullName: "Tester",
		Email: "test@email",
		Password: password,
	}

	err := user.HashPassword()
	require.NoError(t, err)

	require.NotEqual(t, user.Password, password)

	err = user.CheckPassword("wrong password")
	require.Error(t, err)

	err = user.CheckPassword(password)
	require.NoError(t, err)
}
