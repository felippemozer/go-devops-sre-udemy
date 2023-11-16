/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli/utils"

	"github.com/spf13/cobra"
)

// printscreenCmd represents the printscreen command
var printscreenCmd = &cobra.Command{
	Use:   "printscreen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("u")
		utils.GetChromeScreenshot(url, 100)
	},
}

func init() {
	rootCmd.AddCommand(printscreenCmd)
	printscreenCmd.PersistentFlags().String("u", "www.google.com", "URL do screenshot")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printscreenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
