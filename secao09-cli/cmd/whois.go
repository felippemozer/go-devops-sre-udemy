/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli/utils"
	"log"

	"github.com/spf13/cobra"
)

// whoisCmd represents the whois command
var whoisCmd = &cobra.Command{
	Use:   "whois",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("h")
		log.Println(utils.GetWhoIs(host))
	},
}

func init() {
	rootCmd.AddCommand(whoisCmd)
	whoisCmd.PersistentFlags().String("h", "", "Host para verificação do Whois")
}
