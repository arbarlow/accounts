package server

import (
	"github.com/lileio/accounts"
	"github.com/lileio/accounts/database"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func (as AccountsServer) Update(ctx context.Context, r *accounts.UpdateAccountRequest) (*accounts.AccountWithErrorsResponse, error) {
	if r.Account == nil {
		return nil, ErrNoAccount
	}

	a := database.Account{
		ID:    r.Account.Id,
		Name:  r.Account.Name,
		Email: r.Account.Email,
	}

	err := as.DB.Update(&a)
	if err != nil {
		if err == database.ErrAccountNotFound {
			return nil, grpc.Errorf(codes.NotFound, "account not found")
		}
		return nil, err
	}

	res := accounts.AccountWithErrorsResponse{
		Account: accountDetailsFromAccount(&a),
	}

	return &res, nil
}
