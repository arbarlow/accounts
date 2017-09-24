package server

import (
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/lileio/accounts"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestAuthenticate(t *testing.T) {
	truncate()

	ctx := context.Background()
	a := createAccount(t)

	ar := &accounts.AuthenticatePasswordRequest{
		Email:    a.Email,
		Password: pass,
	}

	a, err := s.AuthenticatePassword(ctx, ar)
	assert.NoError(t, err)
	assert.NotEmpty(t, a.Id)
	assert.NotEmpty(t, a.Email)
}

func TestAuthenticateFailure(t *testing.T) {
	truncate()

	ctx := context.Background()
	a := createAccount(t)

	ar := &accounts.AuthenticatePasswordRequest{
		Email:    a.Email,
		Password: "incorrect password lol",
	}

	_, err := s.AuthenticatePassword(ctx, ar)
	assert.Error(t, err)
	assert.Equal(t, grpc.Code(err), codes.PermissionDenied)
}
