package makecmd

import (
	"fmt"
	"go-labs/go-labs/pkg/utils"
	"html/template"
	"os"
	"path"
	"path/filepath"
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

		modelName := utils.Capitalize(args[0])
		withMigration, _ := cmd.Flags().GetBool("migration")

		if err := createModel(modelName); err != nil {
			fmt.Println(" ❌", err)
			return
		}

		if withMigration {
			if err := createMigration(modelName); err != nil {
				fmt.Println(" ❌ ", err)
				return
			}
		}

	},
}

func init() {
	MakeCmd.Flags().BoolP("migration", "m", false, "Create migration file")
}

// ====================================================
// Create Model
// ====================================================
func createModel(modelName string) error {
	modelDir := "internal/domain/models"
	if err := utils.MkdirIfNotExists(modelDir); err != nil {
		return fmt.Errorf("failed to create models directory: %v", err)
	}

	modelFile := filepath.Join(modelDir, fmt.Sprintf("%s.go", strings.ToLower(modelName)))
	if utils.FileIsExists(modelFile) {
		return fmt.Errorf("model %s already exists", modelName)
	}

	file, err := os.Create(modelFile)
	if err != nil {
		return fmt.Errorf("error creating model: %v", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error Closing Model :", err)
		}
	}(file)

	tmpl, _ := template.New("model").Parse(modelTemplate)
	if err := tmpl.Execute(file, map[string]string{"ModelName": modelName}); err != nil {
		return fmt.Errorf("error executing model template: %v", err)
	}

	fmt.Println("Created Model :", path.Base(modelFile))
	return nil
}

// ====================================================
// Create Migration
// ====================================================
func createMigration(modelName string) error {
	migrationDir := "internal/infrastructure/migrations"
	if err := utils.MkdirIfNotExists(migrationDir); err != nil {
		return fmt.Errorf("failed to create migrations directory: %v", err)
	}

	tableName := utils.Pluralize(modelName)
	pattern := fmt.Sprintf("create_%s_table.go", strings.ToLower(tableName))

	files, _ := os.ReadDir(migrationDir)
	fmt.Println(files)
	for _, file := range files {
		if strings.HasSuffix(strings.ToLower(file.Name()), pattern) {
			return fmt.Errorf("migration for \"%s\" already exists", tableName)
		}
	}

	timestamp := time.Now().Format("20060102150405")
	migFile := filepath.Join(migrationDir, fmt.Sprintf("%s_create_%s_table.go", timestamp, tableName))

	file, err := os.Create(migFile)
	if err != nil {
		return fmt.Errorf("error creating migration: %v", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error Closing Migration :", err)
		}
	}(file)

	tmpl, _ := template.New("migration").Parse(migrationTemplate)

	if err := tmpl.Execute(file, map[string]string{"ModelName": modelName}); err != nil {
		return fmt.Errorf("error executing migration template: %v", err)
	}

	fmt.Println("Created Migration :", path.Base(migFile))
	return nil
}
