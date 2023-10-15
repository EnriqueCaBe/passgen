package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmd1Cmd = &cobra.Command{
	Use:   "generate",
	Short: "Descripción del comando 1",
	Run: func(cmd *cobra.Command, args []string) {
		// Lógica para cmd1
		arg1, _ := cmd.Flags().GetString("arg1")
		fmt.Printf("Comando 1 con argumento: %s\n", arg1)
	},
}

func init() {
	rootCmd.AddCommand(cmd1Cmd)
	cmd1Cmd.Flags().String("arg1", "", "Descripción del argumento 1")
}
