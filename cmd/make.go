package cmd

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const modelTemplate = `package models

import "go-labs/go-labs/internal/domain/base"

type {{.ModelName}} struct {
	base.BaseModel
}
`

const migrationTemplate = `package migrations

import (
	"gorm.io/gorm"
)

func Migrate{{.ModelName}}(db *gorm.DB) error {
	type {{.ModelName}} struct {
		
	}

	return db.AutoMigrate(&{{.ModelName}}{})
}
`

var MakeCmd = &cobra.Command{
	Use:   "make:model [name]",
	Short: "Create a new model and migration file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify model name, e.g. artisan make:model Testing -m")
			return
		}

		modelName := strings.Title(args[0])
		withMigration, _ := cmd.Flags().GetBool("migration")

		createModel(modelName)

		if withMigration {
			createMigration(modelName)
		}

	},
}

func init() {
	MakeCmd.Flags().BoolP("migration", "m", false, "Create migration file")
}

func createModel(modelName string) {
	modelDir := "models"
	if _, err := os.Stat(modelDir); os.IsNotExist(err) {
		err := os.MkdirAll(modelDir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating model directory:", err)
			return
		}
	}

	fileName := fmt.Sprintf("%s/%s.go", modelDir, strings.ToLower(modelName))
	tmpl, _ := template.New("model").Parse(modelTemplate)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error Creating Model :", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error Closing Model :", err)
		}
	}(file)

	err = tmpl.Execute(file, map[string]string{
		"ModelName": modelName,
	})

	if err != nil {
		return
	}

	fmt.Println("Created Model :", modelName)
}

func createMigration(modelName string) {
	migrationDir := "migrations"
	if _, err := os.Stat(migrationDir); os.IsNotExist(err) {
		err := os.MkdirAll(migrationDir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating migration directory:", err)
			return
		}
	}

	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("%s/%s_create_%s_table.go", migrationDir, strings.ToLower(modelName), timestamp)
	tmpl, _ := template.New("migration").Parse(migrationTemplate)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error Creating Migration :", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error Closing Migration :", err)
		}
	}(file)

	err = tmpl.Execute(file, map[string]string{
		"ModelName": modelName,
	})
	if err != nil {
		return
	}
	fmt.Println("Created Migration :", modelName)
}
