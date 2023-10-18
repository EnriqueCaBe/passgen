package main

import (
	generatecli "github.com/EnriqueCaBe/passgen/internal/cli"
	"github.com/EnriqueCaBe/passgen/internal/services"
	"github.com/spf13/cobra"
)

func main() {
	generateService := services.NewGenerateService()

	rootCmd := &cobra.Command{Use: "passgen"}
	rootCmd.AddCommand(generatecli.InitGenerateCmd(generateService))
	rootCmd.Execute()
}
