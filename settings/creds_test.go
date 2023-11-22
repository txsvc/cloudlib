package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/txsvc/stdlib/v2"
)

func TestCloneCredentials(t *testing.T) {
	cred := Credentials{
		ProjectID:    "p",
		ClientID:     "c",
		ClientSecret: "s",
		Token:        "t",
		Expires:      10,
		Status:       StateInit,
	}
	dup := cred.Clone()
	assert.Equal(t, &cred, dup)
}

func TestValidation(t *testing.T) {
	cred1 := Credentials{}
	assert.False(t, cred1.IsValid())

	cred2 := Credentials{
		ProjectID:    "p",
		ClientID:     "c",
		ClientSecret: "s",
		Token:        "t",
		Expires:      10, // forces a fail ...
		Status:       StateInit,
	}
	assert.False(t, cred2.IsValid())

	cred2.Expires = 0
	assert.True(t, cred2.IsValid())

	cred2.Status = StateInvalid
	assert.False(t, cred2.IsValid())

	cred2.Status = StateInit // reset

	cred2.Token = ""
	assert.True(t, cred2.IsValid())

	cred2.ClientSecret = ""
	assert.True(t, cred2.IsValid())

	cred2.ClientID = ""
	assert.False(t, cred2.IsValid())

	cred2.Token = "t"
	assert.False(t, cred2.IsValid())

	cred2.Token = ""
	cred2.ClientSecret = "s"
	assert.False(t, cred2.IsValid())
}

func TestExpiration(t *testing.T) {
	cred := Credentials{
		ProjectID:    "p",
		ClientID:     "c",
		ClientSecret: "s",
		Token:        "t",
		Expires:      10,
	}
	assert.True(t, cred.Expired())

	cred.Expires = 0
	assert.False(t, cred.Expired())
}

func TestCredentialsFromEnv(t *testing.T) {
	assert.NotEmpty(t, stdlib.GetString(ClientID, ""))
	assert.NotEmpty(t, stdlib.GetString(ClientSecret, ""))

	cred := CredentialsFromEnv()
	assert.NotNil(t, cred)
	assert.NotEmpty(t, cred)
}
