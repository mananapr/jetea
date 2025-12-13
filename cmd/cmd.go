package cmd

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	cfgFlag string

	rootCmd = &cobra.Command{
		Use:     "jetea",
		Short:   "Jetea is a TUI for NATS.io",
		Long:    `A TUI for NATS.io built with the Bubbletea in Go.`,
		Version: version,
		Args:    cobra.MaximumNArgs(1),
	}
)

func Execute() {
	if err := fang.Execute(context.Background(), rootCmd, fang.WithVersion(version), fang.WithoutCompletions(), fang.WithoutManpage()); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFlag, "config", "c", "", "path to config")
	err := rootCmd.MarkPersistentFlagFilename("config", "yaml", "yml")
	if err != nil {
		log.Fatal("Cannot mark config flag as filename", "error", err)
	}

	rootCmd.Flags().Bool("debug", false, "enable debug mode")

	rootCmd.Run = func(_ *cobra.Command, args []string) {
		_, err := rootCmd.Flags().GetBool("debug")
		if err != nil {
			log.Fatal("Cannot parse debug flag", "error", err)
		}
	}
}
