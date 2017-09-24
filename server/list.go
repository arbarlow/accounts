package server

import (
	"github.com/lileio/accounts"
	context "golang.org/x/net/context"
)

func (as AccountsServer) List(
	ctx context.Context, l *accounts.ListAccountsRequest) (
	*accounts.ListAccountsResponse, error) {

	acs, next_token, err := as.DB.List(l.PageSize, l.PageToken)
	if err != nil {
		return nil, err
	}

	accs := make([]*accounts.Account, len(acs))
	for i, acc := range acs {
		accs[i] = accountDetailsFromAccount(acc)
	}

	return &accounts.ListAccountsResponse{
		Accounts:      accs,
		NextPageToken: next_token,
	}, err
}
