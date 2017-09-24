package server

import (
	"github.com/lileio/accounts"
	"github.com/lileio/accounts/database"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func (as AccountsServer) Create(ctx context.Context, r *accounts.CreateAccountRequest) (*accounts.AccountWithErrorsResponse, error) {
	if r.Account == nil {
		return nil, ErrNoAccount
	}

	a := database.Account{
		Name:     r.Account.Name,
		Email:    r.Account.Email,
		Username: r.Account.Username,
		PhoneNo:  r.Account.PhoneNo,
	}
	err := a.HashPassword(r.Password)
	if err != nil {
		return nil, err
	}

	err = as.DB.Create(&a, r.Password)
	if err != nil {
		if err == database.ErrAccountExists {
			return nil, grpc.Errorf(codes.AlreadyExists, err.Error())
		} else {
			return nil, grpc.Errorf(codes.Internal, err.Error())
		}
	}

	res := accounts.AccountWithErrorsResponse{
		Account: accountDetailsFromAccount(&a),
	}

	return &res, nil
}
