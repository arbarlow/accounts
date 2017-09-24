package server

import (
	"testing"

	"github.com/lileio/accounts"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestGeneratePasswordResetToken(t *testing.T) {
	ctx := context.Background()

	ac := createAccount(t)
	req := &accounts.GeneratePasswordResetTokenRequest{Id: ac.Id}
	res, err := s.GeneratePasswordResetToken(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res.PasswordResetToken)
}
