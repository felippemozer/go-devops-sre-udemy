/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli/utils"

	"github.com/spf13/cobra"
)

// epochConverterCmd represents the epochConverter command
var epochConverterCmd = &cobra.Command{
	Use:   "epochConverter",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		epoch, _ := cmd.Flags().GetString("e")

		if epoch != "" {
			result := utils.ConvertEpoch(epoch)
			cmd.Println(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(epochConverterCmd)
	epochConverterCmd.PersistentFlags().String("e", "", "Tempo em epoch")
}
