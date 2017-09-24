package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/lileio/accounts"
	"github.com/spf13/cobra"
)

var id string

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get an account",
	Run: func(cmd *cobra.Command, args []string) {
		ar := &accounts.GetRequest{
			Id: id,
		}

		ctx := context.Background()
		res, err := client().Get(ctx, ar)
		if err != nil {
			log.Fatal(err)
		}

		js, _ := json.MarshalIndent(res, "", "  ")
		fmt.Println(string(js))
	},
}

func init() {
	RootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&id, "id", "", "", "id (uuid) of account")
}
