package cmd

import (
	"fmt"
	"github.com/ReolinkCameraApi/noctiluca-go-server/internal/pkg/api"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run Reolink Management Server.",
	Long: `Reolink Management Server will run on port 8000 and bind to the host IP. 
	This can be defined with RM_PORT and RM_HOST`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := Serve(); err != nil {
			return err
		}
		return nil
	},
}

func Serve() error {
	router := gin.Default()
	handler := api.NewApiHandler(router)

	handler.Unprotected.GET("/version", func(context *gin.Context) {
		context.String(200, "v0.0.1")
	})

	port := os.Getenv("RM_PORT")
	host := os.Getenv("RM_HOST")

	if port == "" {
		port = "8000"
	}

	if host == "" {
		host = "0.0.0.0"
	}

	log.Printf("Serving on %s:%s", host, port)
	err := endless.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router)
	if err != nil {
		return err
	}
	return nil
}
