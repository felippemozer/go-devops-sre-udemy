/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli/utils"

	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Efetua o encode e decode de strings",
	Long: `Efetua o encode e decode de strings para base64.
	
	Exemplo de uso:
		encode: ./cli base64 --e "string"
		decode: ./cli base64 --d "string"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		encodeStr, _ := cmd.Flags().GetString("e")
		decodeStr, _ := cmd.Flags().GetString("d")

		if encodeStr != "" {
			encode := utils.EncodeString(encodeStr)
			cmd.Println(encode)
		} else if decodeStr != "" {
			decode := utils.DecodeString(decodeStr)
			cmd.Println(decode)
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)

	base64Cmd.PersistentFlags().String("e", "", "Encode da string")
	base64Cmd.PersistentFlags().String("d", "", "Decode da string")
}
