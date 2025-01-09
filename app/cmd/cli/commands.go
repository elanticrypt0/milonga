package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "milonga",
	Short: "Milonga CLI - A tool for managing your Milonga application",
	Long:  `Milonga CLI provides various commands to help you manage and develop your Milonga application.`,
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate various components",
	Long:  `Generate components like models, handlers, and routes.`,
}

var modelCmd = &cobra.Command{
	Use:   "model [name]",
	Short: "Generate a new model with CRUD operations",
	Long: `Generate a new model with its corresponding handler and routes for CRUD operations.
For example:
  milonga generate model User`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		modelName := args[0]
		if err := GenerateModel(modelName); err != nil {
			fmt.Printf("Error generating model: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	generateCmd.AddCommand(modelCmd)
	rootCmd.AddCommand(generateCmd)
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}
