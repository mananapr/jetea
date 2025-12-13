package cmd

import (
	"context"
	"os"
	"time"

	// tea "github.com/charmbracelet/bubbletea"
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
		Long:    "A TUI for NATS.io built with the Bubbletea in Go.",
		Version: version,
		Args:    cobra.MaximumNArgs(1),
	}
)

func Execute() {
	if err := fang.Execute(context.Background(), rootCmd, fang.WithVersion(version), fang.WithoutCompletions(), fang.WithoutManpage()); err != nil {
		os.Exit(1)
	}
}

func initLogger(debug bool) *os.File {
	var logFile *os.File

	log.SetOutput(os.Stderr)
	log.SetLevel(log.FatalLevel)

	if debug {
		newConfigFile, fileErr := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
		if fileErr == nil {
			log.SetTimeFormat(time.Kitchen)
			log.SetLevel(log.DebugLevel)
			log.Info("Logging to debug.log")
			log.SetReportCaller(true)
			log.SetOutput(newConfigFile)
		} else {
			log.Fatal("Unable to open log file", "file", "debug.log", "error", fileErr)
		}
	}

	return logFile
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFlag, "config", "c", "", "path to config")
	err := rootCmd.MarkPersistentFlagFilename("config", "yaml", "yml")
	if err != nil {
		log.Fatal("Cannot mark config flag as filename", "error", err)
	}

	rootCmd.Flags().Bool("debug", false, "enable debug mode")

	rootCmd.Run = func(_ *cobra.Command, args []string) {
		debug, err := rootCmd.Flags().GetBool("debug")
		if err != nil {
			log.Fatal("Cannot parse debug flag", "error", err)
		}
		logFile := initLogger(debug)
		if logFile != nil {
			defer logFile.Close()
		}
	}
}
