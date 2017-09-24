package server

import (
	"testing"

	"github.com/lileio/accounts"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestSearch(t *testing.T) {
	truncate()

	ctx := context.Background()
	a := createAccount(t)
	req := &accounts.SearchRequest{Query: a.Email}

	res, err := cli.Search(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.Accounts), 1)
}
