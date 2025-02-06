package main

import (
	"errors"
	"fmt"
	"net"

	"github.com/jedrw/gowake/cmd/gowake/listen"
	"github.com/jedrw/gowake/pkg/magicpacket"
	"github.com/spf13/cobra"
)

var port int
var ip string

var gowakeCmd = &cobra.Command{
	Use:          "gowake [macaddress]",
	Short:        "Send a magic packet",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ip, _ := cmd.Flags().GetString("ip")
		if net.ParseIP(ip) == nil {
			return errors.New("got invalid IP")
		}

		port, _ := cmd.Flags().GetInt("port")
		magicPacket, err := magicpacket.New(args[0])
		if err != nil {
			return err
		}

		err = magicpacket.Send(magicPacket, ip, port)
		if err != nil {
			return err
		}

		fmt.Printf("Sent magic packet %s to %s:%d\n", args[0], ip, port)
		return nil
	},
}

func init() {
	gowakeCmd.AddCommand(listen.ListenCmd)
	gowakeCmd.Flags().IntVarP(&port, "port", "p", 9, "Port to send magic packet to")
	gowakeCmd.Flags().StringVarP(&ip, "ip", "i", "255.255.255.255", "Destination (IP or broadcast address) for the magic packet")
	gowakeCmd.PersistentFlags().BoolP("help", "h", false, "Print help for command")
	cobra.EnableCommandSorting = false
	gowakeCmd.InitDefaultCompletionCmd()
	gowakeCmd.CompletionOptions.DisableDescriptions = true
}

func main() {
	gowakeCmd.Execute()
}
