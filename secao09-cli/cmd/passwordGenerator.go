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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.Flags().GetInt("s")
		fmt.Println(utils.GeneratePassword(size))
	},
}

func init() {
	rootCmd.AddCommand(passwordGeneratorCmd)
	passwordGeneratorCmd.PersistentFlags().Int("s", 16, "Tamanho da senha")
}
