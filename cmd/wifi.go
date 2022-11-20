package cmd

import (
	"fmt"

	"github.com/AlexEidt/qr"
	"github.com/spf13/cobra"
)

var ssid, password, output string

func init() {
	rootCmd.AddCommand(wifiCmd)
	wifiCmd.Flags().StringVarP(&ssid, "ssid", "s", "", "The ssid of your wifi network")
	wifiCmd.Flags().StringVarP(&password, "password", "p", "", "The password of your wifi network")
	wifiCmd.Flags().StringVarP(&output, "output", "o", "wifi-qr.png", "Specify the location of the generated png")
}

var wifiCmd = &cobra.Command{
	Use:   "wifi",
	Short: "Generates QR for wifi network",
	Long:  `Generates a QR code for the spcified wifi ssid and password`,
	Run: func(cmd *cobra.Command, args []string) {
		qrcode, err := qr.NewQRCode(fmt.Sprintf("WIFI:S:%s;T:WPA;P:%s;;", ssid, password), &qr.Options{Error: "H"})
		if err != nil {
			panic(err)
		}

		qrcode.Render(output, 100)
	},
}
