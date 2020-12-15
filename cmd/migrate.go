package cmd

import (
	"github.com/ReolinkCameraApi/noctiluca-go-server/internal/pkg/database"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run NoctiLuca Database Migration Tool.",
	Long:  `Run NoctiLuca Database Migration Tool will auto-migrate all database structs defined to the set Database.
	To set an external DB use the NOCTI_DB environment variable.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := database.Migrate(); err != nil {
			return err
		}
		return nil
	},
}