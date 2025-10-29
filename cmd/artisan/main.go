package main

import (
	"fmt"
	"os"

	configcmd "go-labs/go-labs/cmd/app/config"
	makecmd "go-labs/go-labs/cmd/app/make"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "artisan",
		Short: "Artisan command line tool for Go Labs",
		Long:  `Artisan is a command line tool that provides various utilities for Go Labs project including model and migration generation.`,
	}

	// Add subcommands
	rootCmd.AddCommand(makecmd.MakeCmd)
	rootCmd.AddCommand(configcmd.ConfigCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
