package listen

import (
	"errors"
	"fmt"
	"syscall"

	"github.com/jedrw/gowake/pkg/magicpacket"
	"github.com/spf13/cobra"
)

var port int
var ip string

var ListenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen for a magic packet",
	RunE: func(cmd *cobra.Command, args []string) error {
		ip, _ := cmd.Flags().GetString("ip")
		port, _ := cmd.Flags().GetInt("port")
		cont, _ := cmd.Flags().GetBool("continuous")
		fmt.Printf("Listening for magic packets on %s:%d\n", ip, port)
		for {
			remote, mac, err := magicpacket.Listen(ip, port)
			if err != nil {
				var errno syscall.Errno
				if errors.As(err, &errno) {
					if errno == syscall.EACCES {
						return fmt.Errorf("%w: please run as elevated user", err)
					}
				} else {
					return err
				}
			}

			fmt.Printf("%s from %s\n", mac, remote.String())
			if !cont {
				break
			}
		}

		return nil
	},
}

func init() {
	var continuous bool
	ListenCmd.Flags().IntVarP(&port, "port", "p", 9, "Port to listen for magic packets on")
	ListenCmd.Flags().StringVarP(&ip, "ip", "i", "0.0.0.0", "Address to listen for magic packets on")
	ListenCmd.Flags().BoolVarP(&continuous, "continuous", "c", false, "Listen continuously for magic packets")
}
