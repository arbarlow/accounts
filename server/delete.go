package server

import (
	"github.com/lileio/accounts"
	"github.com/lileio/accounts/database"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func (as AccountsServer) Delete(ctx context.Context, r *accounts.DeleteAccountRequest) (*accounts.Account, error) {
	ca, err := as.DB.Get(r.Id, "")
	if err != nil {
		if err == database.ErrAccountNotFound {
			return nil, grpc.Errorf(codes.NotFound, "account not found")
		}
		return nil, err
	}

	err = as.DB.Delete(r.Id)
	if err != nil {
		return nil, err
	}

	return accountDetailsFromAccount(ca), nil
}
