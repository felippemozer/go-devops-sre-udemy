/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli/utils"

	"github.com/spf13/cobra"
)

// portCheckerCmd represents the portChecker command
var portCheckerCmd = &cobra.Command{
	Use:   "portChecker",
	Short: "Verificar quais portas estão abertas para um host",
	Long: `Checa uma lista de portas de um host para verificar
	quais portas estão abertas para acesso e quais estão fechadas.
	
	Exemplo de uso:
		./cli portChecker --h www.google.com
		./cli portChecker --h www.google.com --p 80,443,22,56
	`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("h")
		ports, _ := cmd.Flags().GetString("p")

		utils.CheckPort(host, ports)
	},
}

func init() {
	rootCmd.AddCommand(portCheckerCmd)

	portCheckerCmd.PersistentFlags().String("h", "", "Host a ser validado")
	portCheckerCmd.PersistentFlags().String("p", "443", "Lista de portas separada por vírgula. Ex: 80,443,22")
}
