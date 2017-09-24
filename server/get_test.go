package server

import (
	"testing"

	"github.com/lileio/accounts"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestGet(t *testing.T) {
	truncate()

	ctx := context.Background()
	a := createAccount(t)
	req := &accounts.GetRequest{Id: a.Id}

	res, err := cli.Get(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res.Id)
}
