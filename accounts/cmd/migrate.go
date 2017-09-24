package cmd

import (
	"github.com/lileio/accounts/database"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		conn := database.DatabaseFromEnv()
		conn.Migrate()
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}
