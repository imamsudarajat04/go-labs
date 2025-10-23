package main

import (
	"fmt"
	"go-labs/go-labs/cmd"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "artisan",
	}
	rootCmd.AddCommand(cmd.MakeCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
