package main

import (
	"fmt"
	"go-labs/go-labs/internal/config"
)

func main() {
	fmt.Printf("🚀 Starting Go Labs Application... \n")

	// Load configuration
	fmt.Println("📁 Loading configuration...")
	config.LoadConfig()
}
