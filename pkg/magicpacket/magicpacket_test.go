package magicpacket_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/jedrw/gowake/pkg/magicpacket"
)

func TestNewErrorsWithInvalidMacAddress(t *testing.T) {
	_, err := magicpacket.New("ab:ab:ab:ab:ab:ag")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestNewErrorsWithNonEUI48MacAddress(t *testing.T) {
	_, err := magicpacket.New("ab:ab:ab:ab:ab:ab:ab:ab")
	if !errors.Is(err, magicpacket.ErrNotValidEUI48MacAddress) {
		t.Errorf("expected %s, Got: %s", magicpacket.ErrNotValidEUI48MacAddress, err)
	}
}

func TestNewReturnsValidMagicPacket(t *testing.T) {
	magicPacket, err := magicpacket.New("ab:ab:ab:ab:ab:ab")
	if err != nil {
		t.Error(err)
	}

	expectedBytes := []byte{255, 255, 255, 255, 255, 255, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171}
	if !bytes.Equal(magicPacket.Bytes(), expectedBytes) {
		t.Errorf("expected: %b, got: %b", expectedBytes, magicPacket.Bytes())
	}
}

func TestMagicPacketMac(t *testing.T) {
	magicPacket, err := magicpacket.New("ab:ab:ab:ab:ab:ab")
	if err != nil {
		t.Error(err)
	}

	expectedMac := "ab:ab:ab:ab:ab:ab"
	if magicPacket.Mac() != expectedMac {
		t.Errorf("expected: %s, got: %s", expectedMac, magicPacket.Mac())
	}
}

func TestMagicPacketValidateReturnsErrorOnMalformedMagicPacket(t *testing.T) {
	magicPacket := magicpacket.MagicPacket{255, 255, 255, 255, 255, 255}
	offset := 6
	for i := 0; i < 15; i++ {
		copy(magicPacket[offset:], []byte{171, 171, 171, 171, 171, 171})
		offset += 6
	}

	copy(magicPacket[offset:], []byte{172, 172, 172, 172, 172, 172})

	err := magicPacket.Validate()
	if !errors.Is(err, magicpacket.ErrMalformedMagicPacket) {
		t.Errorf("expected %s, Got: %s", magicpacket.ErrMalformedMagicPacket, err)
	}
}
