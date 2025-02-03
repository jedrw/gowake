package magicpacket

import (
	"bytes"
	"fmt"
	"net"
)

func Listen(port int) (*net.UDPAddr, string, error) {
	addr := net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: port,
	}

	listener, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return nil, "", err
	}
	defer listener.Close()

	var magicPacket MagicPacket
	remote := &net.UDPAddr{}
	_, remote, err = listener.ReadFromUDP(magicPacket[:])
	if err != nil {
		return remote, "", err
	}

	macLength := 6
	offset := 6
	for i := 0; i < 16; i++ {
		if !bytes.Equal(magicPacket[offset:offset+macLength], magicPacket[96:]) {
			return remote, "", fmt.Errorf("received malformed magicpacket from %v", remote)
		}
		offset += 6
	}

	return remote, net.HardwareAddr.String(magicPacket[96:]), err
}
