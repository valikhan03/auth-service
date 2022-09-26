package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGenerateAndValidateToken(t *testing.T) {
	wrapper := JWTWrapper{
		SecretKey: "test-secret-key",
		Issuer: "test-issuer",
		ExpirationTime: 100 * time.Second,
	}

	test_user := User{
		Email: "test@email.com",
		ID: 4654115,
		FullName: "Zhaqsylyq Zhaqsylyq",
	}

	token, err := wrapper.GenerateToken(test_user)
	require.NoError(t, err)

	require.Greater(t, len(token), 0)

	claims, err := wrapper.ValidateToken(token)
	require.NoError(t, err)

	require.NotNil(t, claims)

	require.Equal(t, test_user.Email, claims.Email)
	require.Equal(t, test_user.ID, claims.ID)
	require.Equal(t, test_user.FullName, claims.FullName)
}

func TestExpiredToken(t *testing.T) {
	wrapper := JWTWrapper{
		SecretKey: "test-secret-key",
		Issuer: "test-issuer",
		ExpirationTime: 10 * time.Second,
	}

	test_user := User{
		Email: "test@email.com",
		ID: 4654115,
		FullName: "Zhaqsylyq Zhaqsylyq",
	}

	token, err := wrapper.GenerateToken(test_user)
	require.NoError(t, err)

	time.Sleep(11 * time.Second)

	_, err = wrapper.ValidateToken(token)
	require.Error(t, err)
}