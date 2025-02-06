package magicpacket

import (
	"fmt"
	"net"
)

func Listen(ip string, port int) (*net.UDPAddr, string, error) {
	listenIP := net.ParseIP(ip)
	if listenIP == nil {
		return nil, "", fmt.Errorf("invalid IP: %s", ip)
	}

	addr := net.UDPAddr{
		IP:   listenIP,
		Port: port,
	}

	listener, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return nil, "", err
	}
	defer listener.Close()

	magicPacket := MagicPacket{}
	_, remote, err := listener.ReadFromUDP(magicPacket[:])
	if err != nil {
		return remote, "", err
	}

	return remote, magicPacket.Mac(), magicPacket.Validate()
}
