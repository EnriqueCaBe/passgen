package cmd

import (
	"errors"

	"github.com/EnriqueCaBe/passgen/internal/services"
	"github.com/spf13/cobra"
)

var length int8
var hasLetters, hasSymbols, hasNumbers bool

type CobraFn func(cmd *cobra.Command, args []string)

func InitGenerateCmd(service services.GenerateService) *cobra.Command {
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generates a password",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if length < 4 || length > 40 {
				return errors.New("length must be between 4 and 40")
			}

			if !hasLetters && !hasSymbols && !hasNumbers {
				return errors.New("you must include at least one type of character in the password")
			}

			return nil
		},
		Run: runGenerateFn(service),
	}

	generateCmd.PersistentFlags().Int8VarP(&length, "length", "l", 0, "Password length")
	generateCmd.PersistentFlags().BoolVarP(&hasLetters, "letters", "L", false, "Include letters")
	generateCmd.PersistentFlags().BoolVarP(&hasSymbols, "symbols", "S", false, "Include symbols")
	generateCmd.PersistentFlags().BoolVarP(&hasNumbers, "numbers", "N", false, "Include numbers")

	return generateCmd
}

func runGenerateFn(service services.GenerateService) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		service.GeneratePassword(&services.GenerateCommand{
			Length:  length,
			Letters: hasLetters,
			Symbols: hasSymbols,
			Numbers: hasNumbers,
		})
	}
}
