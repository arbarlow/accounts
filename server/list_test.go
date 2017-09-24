package server

import (
	"os"
	"testing"

	"github.com/lileio/accounts"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestSimpleList(t *testing.T) {
	truncate()

	for i := 0; i < 5; i++ {
		createAccount(t)
	}

	ctx := context.Background()
	req := &accounts.ListAccountsRequest{
		PageSize: 6,
	}

	l, err := s.List(ctx, req)
	assert.Nil(t, err)
	assert.Empty(t, l.NextPageToken)
}

func TestSimpleListToken(t *testing.T) {
	if os.Getenv("CASSANDRA_DB_NAME") != "" {
		t.Skip()
	}

	truncate()

	acc := []*accounts.Account{}
	for i := 0; i < 4; i++ {
		acc = append(acc, createAccount(t))
	}

	ctx := context.Background()
	req := &accounts.ListAccountsRequest{
		PageSize: 2,
	}

	l, err := s.List(ctx, req)
	assert.Nil(t, err)
	assert.NotEmpty(t, l.NextPageToken)

	req = &accounts.ListAccountsRequest{
		PageSize:  2,
		PageToken: l.NextPageToken,
	}

	l, err = s.List(ctx, req)
	assert.Nil(t, err)
	assert.Empty(t, l.NextPageToken)
}
