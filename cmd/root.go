package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qrgen",
	Short: "QR code generator",
	Long:  `Simple qr code generation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("here")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
