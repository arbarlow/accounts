package server

import (
	"context"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/lileio/accounts"
	"github.com/lileio/accounts/database"
	"github.com/lileio/lile"
	"github.com/stretchr/testify/assert"
)

var s = AccountsServer{}
var db = setupDB()
var cli accounts.AccountsClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		accounts.RegisterAccountsServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = accounts.NewAccountsClient(lile.TestConn(addr))

	os.Exit(m.Run())
}

func setupDB() database.Database {
	conn := database.DatabaseFromEnv()
	conn.(*database.PostgreSQL).DB.Exec("drop table accounts;")
	s.DB = conn
	s.DB.Migrate()
	return conn
}

func truncate() {
	db.Truncate()
}

var name = "Alex B"
var email = "alexb@localhost"
var pass = "password"

func createAccount(t *testing.T) *accounts.Account {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ctx := context.Background()
	req := &accounts.CreateAccountRequest{
		Account: &accounts.Account{
			Name:  name,
			Email: email + strconv.Itoa(r.Int()),
		},
		Password: pass,
	}
	account, err := s.Create(ctx, req)
	assert.NoError(t, err)
	return account.Account
}
