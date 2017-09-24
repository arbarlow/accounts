package server

import (
	"github.com/lileio/accounts"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func (as AccountsServer) AuthenticatePassword(ctx context.Context, r *accounts.AuthenticatePasswordRequest) (*accounts.Account, error) {
	a, err := as.DB.Get(r.Id, r.Email)
	if err != nil {
		return nil, err
	}

	err = a.ComparePasswordToHash(r.Password)
	if err != nil {
		return nil, grpc.Errorf(codes.PermissionDenied, "password incorrect")
	}

	return accountDetailsFromAccount(a), nil
}
