/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// passwordGeneratorCmd represents the passwordGenerator command
var passwordGeneratorCmd = &cobra.Command{
	Use:   "passwordGenerator",
	Short: "Gerador de senhas",
	Long: `Um gerador de senhas pseudorandomicas. Necessário um mínimo de 8
	caracteres por senha gerada.

	Exemplos de uso:
		./cli passwordGenerator (sem opção: gera senhas com 16 caracteres)
		./cli passwordGenerator --s 32
	`,
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.Flags().GetInt("s")
		fmt.Println(utils.GeneratePassword(size))
	},
}

func init() {
	rootCmd.AddCommand(passwordGeneratorCmd)
	passwordGeneratorCmd.PersistentFlags().Int("s", 16, "Tamanho da senha (Mínimo de 8 caracteres)")
}
