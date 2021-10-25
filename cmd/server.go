package cmd

import (
	"github.com/barqus/fillq_backend/internal/handler"
	"github.com/spf13/cobra"
	"os"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		handler.HandleHTTP()
		os.Exit(0)
	},
}