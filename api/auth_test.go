package api_test

import (
	"os"
	"testing"

	"github.com/SevereCloud/vksdk/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_AuthCheckPhone(t *testing.T) {
	needServiceToken(t)

	clientSecret := os.Getenv("CLIENT_SECRET")
	clientID := os.Getenv("CLIENT_ID")

	if clientSecret == "" || clientID == "" {
		t.Skip("need params")
	}

	_, err := vkUser.AuthCheckPhone(api.Params{
		"phone":         "+79523071234",
		"client_id":     clientID,
		"client_secret": clientSecret,
	})
	assert.NoError(t, err)
}
