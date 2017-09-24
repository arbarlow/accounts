package server

import (
	"github.com/lileio/accounts"
	"github.com/lileio/accounts/database"
	"github.com/lileio/image_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type AccountsServer struct {
	accounts.AccountsServer
	DB database.Database
}

var (
	is image_service.ImageServiceClient

	ErrNoAccount = grpc.Errorf(codes.InvalidArgument, "account is nil/no account provided")
)

func accountDetailsFromAccount(a *database.Account) *accounts.Account {
	return &accounts.Account{
		Id:                 a.ID,
		Name:               a.Name,
		Email:              a.Email,
		ConfirmToken:       a.ConfirmationToken,
		PasswordResetToken: a.PasswordResetToken,
	}
}
