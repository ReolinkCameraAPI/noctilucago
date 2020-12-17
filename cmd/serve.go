package cmd

import (
	"fmt"
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/api"
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/procedures"
	"github.com/fvbock/endless"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run NoctiLuca Server.",
	Long: `NoctiLuca Server will run on port 8000 and bind to the host IP. 
	This can be defined with NOC_PORT and NOC_HOST`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := Serve(); err != nil {
			return err
		}
		return nil
	},
}

func Serve() error {
	db, err := procedures.NewDatabase()

	if err != nil {
		return err
	}

	handler := api.NewApiHandler(db)
	handler.CreateEndpoints()

	port := os.Getenv("NOC_PORT")
	host := os.Getenv("NOC_HOST")

	if port == "" {
		port = "8000"
	}

	if host == "" {
		host = "0.0.0.0"
	}

	log.Printf("Serving on %s:%s", host, port)

	if err := endless.ListenAndServe(fmt.Sprintf("%s:%s", host, port), handler.Router); err != nil {
		return err
	}

	return nil
}
