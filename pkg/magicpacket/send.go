package magicpacket

import (
	"fmt"
	"net"
)

func Send(magicPacket MagicPacket, ip string, port int) error {
	sendIP := net.ParseIP(ip)
	if sendIP == nil {
		return fmt.Errorf("invalid IP: %s", ip)
	}

	addr := net.UDPAddr{
		IP:   sendIP,
		Port: port,
	}

	conn, err := net.Dial("udp", addr.String())
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(magicPacket.Bytes())
	if err != nil {
		return err
	}

	return err
}
