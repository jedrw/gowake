package magicpacket

import (
	"fmt"
	"net"
)

type MagicPacket [102]byte

func New(mac string) (MagicPacket, error) {
	// Parse mac address
	hwAddr, err := net.ParseMAC(mac)
	if err != nil {
		return MagicPacket{}, err
	}

	if len(hwAddr) != 6 {
		return MagicPacket{}, fmt.Errorf("invalid EUI-48 MAC address")
	}

	// Build magicpacket
	magicPacket := MagicPacket{255, 255, 255, 255, 255, 255}

	offset := 6
	for i := 0; i < 16; i++ {
		copy(magicPacket[offset:], hwAddr[:])
		offset += 6
	}

	return magicPacket, err
}
