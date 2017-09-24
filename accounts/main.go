package main

import (
	"github.com/lileio/accounts"
	"github.com/lileio/accounts/accounts/cmd"
	"github.com/lileio/accounts/database"
	"github.com/lileio/accounts/server"
	"github.com/lileio/lile"
	"google.golang.org/grpc"
)

func main() {
	db := database.DatabaseFromEnv()
	db.Migrate()
	defer db.Close()

	s := &server.AccountsServer{DB: db}

	lile.Name("accounts")
	lile.Server(func(g *grpc.Server) {
		accounts.RegisterAccountsServer(g, s)
	})

	lile.AddPubSubInterceptor(map[string]string{
		"Create":                     "account_service.created",
		"Update":                     "account_service.updated",
		"Delete":                     "account_service.deleted",
		"GeneratePasswordResetToken": "account_service.generated_password_reset",
	})

	cmd.Execute()
}
