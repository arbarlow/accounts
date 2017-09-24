package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/lileio/accounts"
	"github.com/spf13/cobra"
)

var cfgFile string
var addr string

var client = func() accounts.AccountsClient {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithTimeout(1*time.Second),
	)

	if err != nil {
		log.Fatal(err)
	}

	return accounts.NewAccountsClient(conn)
}

var RootCmd = &cobra.Command{
	Use:   "accounts_cli",
	Short: "A cli for accounts",
}

func Execute() {
	RootCmd.PersistentFlags().StringVarP(&addr, "addr", "a", "localhost:8000", "address for service. i.e localhost:8001")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
