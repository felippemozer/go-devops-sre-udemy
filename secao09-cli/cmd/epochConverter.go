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
	Short: "Converte um UNIX epoch time para Data",
	Long: `Converte uma entrada no formato UNIX epoch para o formato
	de Data. 
	
	Exemplo de uso;
		./cli epochConverter --e 1700094502
	`,
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
