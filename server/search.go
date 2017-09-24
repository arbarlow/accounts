package server

import (
	"github.com/lileio/accounts"
	context "golang.org/x/net/context"
)

func (s AccountsServer) Search(ctx context.Context, r *accounts.SearchRequest) (*accounts.SearchResponse, error) {
	acs, err := s.DB.Search(r.Query)
	if err != nil {
		return nil, err
	}

	accs := make([]*accounts.Account, len(acs))
	for i, acc := range acs {
		accs[i] = accountDetailsFromAccount(acc)
	}

	return &accounts.SearchResponse{Accounts: accs}, nil
}
