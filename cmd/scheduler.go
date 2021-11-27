package cmd

import (
	"github.com/barqus/fillq_backend/internal/scheduler"
	"github.com/spf13/cobra"
	"os"
)

var schedulerCmd = &cobra.Command{
	Use: "scheduler",
	Run: func(cmd *cobra.Command, args []string) {
		scheduler.RunTasks()
		os.Exit(0)
	},
}

