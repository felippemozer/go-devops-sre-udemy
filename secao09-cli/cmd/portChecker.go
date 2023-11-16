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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("h")
		ports, _ := cmd.Flags().GetString("p")

		utils.CheckPort(host, ports)
	},
}

func init() {
	rootCmd.AddCommand(portCheckerCmd)

	portCheckerCmd.PersistentFlags().String("h", "", "Host a ser validado")
	portCheckerCmd.PersistentFlags().String("p", "", "Lista de portas separada por vírgula. Ex: 80,443,22")
}
