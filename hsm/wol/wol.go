package wol

import (
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const pckt_len = 6 + 16*6

type MacAddress [6]byte
type MagicPacket [pckt_len]byte

// Creates a new MacAddress from string.
// Returns an error in case of failing to parse the string.
func NewMacAddressFromString(addrStr string) (MacAddress, error) {
	match, err := regexp.MatchString("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$", addrStr)
	if err != nil {
		return MacAddress{}, err
	}

	if match == false {
		return MacAddress{}, errors.New(fmt.Sprintf("NewMacAddressFromString: addrStr does not match regexp: \"%s\"", addrStr))
	}

	addrStr = strings.ReplaceAll(addrStr, ":", "")
	addrStr = strings.ReplaceAll(addrStr, "-", "")

	res, err := hex.DecodeString(addrStr)
	if err != nil {
		return MacAddress{}, err
	}

	return MacAddress{res[0], res[1], res[2], res[3], res[4], res[5]}, nil
}

// Creates a new Packet from a MacAddress.
func NewPacket(macAddress MacAddress) MagicPacket {
	var packet MagicPacket

	packet[0] = 0xFF
	packet[1] = 0xFF
	packet[2] = 0xFF
	packet[3] = 0xFF
	packet[4] = 0xFF
	packet[5] = 0xFF

	for i := 6; i < pckt_len; i++ {
		packet[i] = macAddress[i%6]
	}

	return packet
}

// Is equivalent to calling NewMacAddressFromString() and NewPacket() subsequently.
func NewPacketFromString(addrStr string) (MagicPacket, error) {
	addr, err := NewMacAddressFromString(addrStr)
	if err != nil {
		return MagicPacket{}, err
	}

	packet := NewPacket(addr)
	return packet, nil
}

func BroadcastPacket(magicPacket MagicPacket) error {

	return nil
}
