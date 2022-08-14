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

	test_email := "test@email.com"

	token, err := wrapper.GenerateToken(test_email)
	require.NoError(t, err)

	require.Greater(t, len(token), 0)

	claims, err := wrapper.ValidateToken(token)
	require.NoError(t, err)

	require.NotNil(t, claims)

	require.Equal(t, test_email, claims.Email)
}

func TestExpiredToken(t *testing.T) {
	wrapper := JWTWrapper{
		SecretKey: "test-secret-key",
		Issuer: "test-issuer",
		ExpirationTime: 10 * time.Second,
	}

	test_email := "test@email.com"

	token, err := wrapper.GenerateToken(test_email)
	require.NoError(t, err)

	time.Sleep(11 * time.Second)

	_, err = wrapper.ValidateToken(token)
	require.Error(t, err)
}