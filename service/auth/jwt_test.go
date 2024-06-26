package auth

import (
	"testing"
)

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJWT(secret, 1)
	if err != nil {
		t.Errorf("Error creating JWT: %v", err)
	} 
	if token == "" {
		t.Error("Expected token to be not empty")
	}
}