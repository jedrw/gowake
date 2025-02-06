package magicpacket

import (
	"bytes"
	"errors"
	"fmt"
	"net"
)

var (
	ErrNotValidEUI48MacAddress = errors.New("not a valid EUI-48 MAC address")
	ErrMalformedMagicPacket    = errors.New("malformed magic packet")
)

type MagicPacket [102]byte

func New(mac string) (MagicPacket, error) {
	hwAddr, err := net.ParseMAC(mac)
	if err != nil {
		return MagicPacket{}, err
	}

	if len(hwAddr) != 6 {
		return MagicPacket{}, fmt.Errorf("%w, %s", ErrNotValidEUI48MacAddress, hwAddr)
	}

	magicPacket := MagicPacket{255, 255, 255, 255, 255, 255}
	offset := 6
	for i := 0; i < 16; i++ {
		copy(magicPacket[offset:], hwAddr[:])
		offset += 6
	}

	return magicPacket, nil
}

func (mp *MagicPacket) Bytes() []byte {
	return mp[:]
}

func (mp MagicPacket) Mac() string {
	return net.HardwareAddr.String(mp[96:])
}

func (mp MagicPacket) Validate() error {
	macLength := 6
	offset := 6
	for i := 0; i < 16; i++ {
		if !bytes.Equal(mp[offset:offset+macLength], mp[96:]) {
			return ErrMalformedMagicPacket
		}

		offset += 6
	}

	_, err := net.ParseMAC(mp.Mac())
	if err != nil {
		return err
	}

	return nil
}
