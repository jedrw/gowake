package listen

import (
	"fmt"

	"github.com/jedrw/gowake/pkg/magicpacket"
	"github.com/spf13/cobra"
)

var ListenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen for a magic packet",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		cont, _ := cmd.Flags().GetBool("continuous")
		fmt.Printf("Listening for magic packets on port %d:\n", port)
		for {
			remote, mac, err := magicpacket.Listen(port)
			if err != nil {
				if err.Error() == fmt.Sprintf("listen udp 0.0.0.0:%d: bind: permission denied", port) {
					fmt.Println("Please run as elevated user")
					return
				} else {
					fmt.Println(err)
					return
				}
			}
			fmt.Printf("%v from %v\n", mac, remote.String())
			if !cont {
				break
			}
		}
	},
}

func init() {
	var continuous bool
	ListenCmd.Flags().BoolVarP(&continuous, "continuous", "c", false, "Listen continuously for magic packets")
}
