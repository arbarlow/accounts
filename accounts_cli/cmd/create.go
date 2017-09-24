package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/lileio/accounts"
	"github.com/spf13/cobra"
)

var name string
var email string
var password string
var image string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create an account",
	Run: func(cmd *cobra.Command, args []string) {
		ar := &accounts.CreateAccountRequest{
			Account: &accounts.Account{
				Name:  name,
				Email: email,
			},
			Password: password,
		}

		ctx := context.Background()
		res, err := client().Create(ctx, ar)
		if err != nil {
			log.Fatal(err)
		}

		js, _ := json.MarshalIndent(res, "", "  ")
		fmt.Println(string(js))
	},
}

func init() {
	createCmd.Flags().StringVarP(&name, "name", "n", "", "name for account")
	createCmd.Flags().StringVarP(&email, "email", "e", "", "email address")
	createCmd.Flags().StringVarP(&password, "password", "p", "", "password for account")

	RootCmd.AddCommand(createCmd)
}
