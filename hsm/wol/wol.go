package wol

const pckt_len = 6 + 16*6

type MacAddress [6]byte
type MagicPacket [pckt_len]byte

func MakeWolPacket(macAddress MacAddress) MagicPacket {
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

func BroadcastWolPacket(magicPacket MagicPacket) error {

	return nil
}
