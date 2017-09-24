package server

import (
	"github.com/lileio/accounts"
	"github.com/lileio/accounts/database"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func (s AccountsServer) GeneratePasswordResetToken(ctx context.Context, r *accounts.GeneratePasswordResetTokenRequest) (*accounts.Account, error) {
	a, err := s.DB.GeneratePasswordToken(r.Id, r.Email)
	if err != nil && err == database.ErrAccountNotFound {
		return nil, grpc.Errorf(codes.NotFound, "account not found")
	}

	if err != nil {
		return nil, err
	}

	return accountDetailsFromAccount(a), nil
}
