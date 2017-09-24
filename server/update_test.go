package server

import (
	"testing"

	"github.com/lileio/accounts"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestUpdateSuccess(t *testing.T) {
	truncate()

	ctx := context.Background()
	a := createAccount(t)

	email := "somethingnew@gmail.com"
	a.Email = email

	ar := &accounts.UpdateAccountRequest{
		Account: a,
	}

	a2, err := s.Update(ctx, ar)
	assert.Nil(t, err)
	assert.NotNil(t, a2)
	assert.NotEmpty(t, a2.Account.Id)
	assert.Equal(t, a2.Account.Email, email)

	a3, err := s.Get(ctx, &accounts.GetRequest{Id: a.Id})
	assert.Nil(t, err)
	assert.Equal(t, a3.Email, email)
}

func TestUpdateNotExist(t *testing.T) {
	truncate()
	ctx := context.Background()
	// u1 := uuid.NewV1()

	ar := &accounts.UpdateAccountRequest{
		Account: nil,
	}

	a2, err := s.Update(ctx, ar)
	assert.NotNil(t, err)
	assert.Nil(t, a2)
}
