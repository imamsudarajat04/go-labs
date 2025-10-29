package configcmd

import (
	"fmt"
	"os"

	"go-labs/go-labs/internal/config"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration management commands",
	Long:  `Configuration management commands for Go Labs application`,
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check and validate configuration",
	Long:  `Check and validate the current configuration from .env file and environment variables`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔍 Checking configuration...")
		fmt.Println("")

		// Load configuration with strict validation
		cfg, err := config.LoadStrict()
		if err != nil {
			fmt.Printf("❌ Failed to load configuration: %v\n", err)
			fmt.Println("")
			fmt.Println("💡 Please check your .env file or environment variables.")
			fmt.Println("   You can copy env.example to .env and update the values.")
			os.Exit(1)
		}

		// Validate configuration
		if err := cfg.Validate(); err != nil {
			fmt.Printf("❌ Configuration validation failed:\n%v\n", err)
			fmt.Println("")
			fmt.Println("💡 Please check your .env file or environment variables.")
			fmt.Println("   You can copy env.example to .env and update the values.")
			os.Exit(1)
		}

		// Print configuration
		fmt.Println("✅ Configuration is valid!")
		fmt.Println("")
		cfg.Print()
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Long:  `Show the current configuration without validation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📋 Current Configuration:")
		fmt.Println("")

		// Load configuration
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("❌ Failed to load configuration: %v\n", err)
			os.Exit(1)
		}

		// Print configuration
		cfg.Print()
	},
}

func init() {
	ConfigCmd.AddCommand(checkCmd)
	ConfigCmd.AddCommand(showCmd)
}
