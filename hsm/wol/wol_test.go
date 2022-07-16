package wol

import "testing"

func TestNewMacAddressFromString(t *testing.T) {
	addrStr1 := "01-02-05-FF-FE-55"
	addrStr2 := "01-02-05-ff-fe-55"
	addrStr3 := "01:02:05:FF:FE:55"
	addrStr4 := "01.02.05.FF.FE.55"
	addrStr5 := "01-02-05-FF-GE-55"

	// addrStr1
	addr, err := NewMacAddressFromString(addrStr1)
	if err != nil {
		t.Errorf("%s \"%s\"", err.Error(), addrStr1)
	}

	if !(addr[0] == 0x01 && addr[1] == 0x02 && addr[2] == 0x05 && addr[3] == 0xFF && addr[4] == 0xFE && addr[5] == 0x55) {
		t.Errorf("Unexpected values: \"%s\"", addrStr1)
	}

	// addrStr2
	addr, err = NewMacAddressFromString(addrStr2)
	if err != nil {
		t.Errorf("%s \"%s\"", err.Error(), addrStr2)
	}

	if !(addr[0] == 0x01 && addr[1] == 0x02 && addr[2] == 0x05 && addr[3] == 0xFF && addr[4] == 0xFE && addr[5] == 0x55) {
		t.Errorf("Unexpected values: \"%s\"", addrStr2)
	}

	// addrStr3
	addr, err = NewMacAddressFromString(addrStr3)
	if err != nil {
		t.Errorf("%s \"%s\"", err.Error(), addrStr3)
	}

	if !(addr[0] == 0x01 && addr[1] == 0x02 && addr[2] == 0x05 && addr[3] == 0xFF && addr[4] == 0xFE && addr[5] == 0x55) {
		t.Errorf("Unexpected values: \"%s\"", addrStr3)
	}

	// addrStr4
	addr, err = NewMacAddressFromString(addrStr4)
	if err == nil {
		t.Errorf("Should return an error: \"%s\"", addrStr4)
	}

	// addrStr5
	addr, err = NewMacAddressFromString(addrStr5)
	if err == nil {
		t.Errorf("Should return an error: \"%s\"", addrStr5)
	}
}

func TestNewPacket(t *testing.T) {
	address := MacAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}
	packet := NewPacket(address)

	for i := 0; i < 6; i++ {
		if !(packet[i] == 0xFF) {
			t.Errorf("Malformed packet. {i = %d}", i)
			break
		}
	}

	for i := 6; i < len(packet); i++ {
		if !(packet[i] == address[i%6]) {
			t.Errorf("Malformed packet. {i = %d}", i)
			break
		}
	}
}

func TestNewPacketFromString(t *testing.T) {
	addressCheck := MacAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}
	address := "01:02:03:04:05:06"
	packet, err := NewPacketFromString(address)
	if err != nil {
		t.Errorf(err.Error())
	}

	for i := 0; i < 6; i++ {
		if !(packet[i] == 0xFF) {
			t.Errorf("Malformed packet. {i = %d}", i)
			break
		}
	}

	for i := 6; i < len(packet); i++ {
		if !(packet[i] == addressCheck[i%6]) {
			t.Errorf("Malformed packet. {i = %d}", i)
			break
		}
	}
}

func TestSendPacket(t *testing.T) {
	packet, err := NewPacketFromString("00:00:00:00:00:00")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = SendPacket(packet, "127.0.0.1")
	if err != nil {
		t.Errorf(err.Error())
	}
}
