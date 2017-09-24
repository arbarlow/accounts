package server

import (
	"context"
	"strconv"
	"testing"

	"github.com/lileio/accounts"
	"github.com/stretchr/testify/assert"
)

func TestCreateSuccess(t *testing.T) {
	truncate()

	account := createAccount(t)
	assert.NotEmpty(t, account.Id)
}

func BenchmarkCreate(b *testing.B) {
	truncate()

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		req := &accounts.CreateAccountRequest{
			Account: &accounts.Account{
				Name:  name,
				Email: "alexbarlowis@localhost" + strconv.Itoa(i),
			},
			Password: pass,
		}

		_, err := s.Create(ctx, req)

		if err != nil {
			panic(err)
		}
	}
}

func TestCreateUniqueness(t *testing.T) {
	truncate()

	ctx := context.Background()
	a1 := createAccount(t)

	req2 := &accounts.CreateAccountRequest{
		Account:  a1,
		Password: pass,
	}

	a2, err := s.Create(ctx, req2)
	assert.Error(t, err)
	assert.Nil(t, a2)
}

func TestCreateEmpty(t *testing.T) {
	truncate()

	ctx := context.Background()
	req := &accounts.CreateAccountRequest{}

	account, err := s.Create(ctx, req)
	assert.NotNil(t, err)
	assert.Nil(t, account)
}
