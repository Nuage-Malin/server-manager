package hsm

import (
	"nuagemalin.com/hsm/hsmlib/conf"
	"nuagemalin.com/hsm/hsmlib/sshc"
	"nuagemalin.com/hsm/hsmlib/wol"
)

//export LoadConf
func LoadConf(file string) *conf.ConfFile {
	c, err := conf.Load(file)
	if err != nil {
		return nil
	}

	return c
}

func Wake(conf *conf.ConfFile, machine string) error {
	serv, err := conf.FindServerUnitByName(machine)
	if err != nil {
		return err
	}

	packet, err := wol.NewPacketFromString(serv.MacAddress)
	if err != nil {
		return err
	}

	err = wol.SendPacket(packet, serv.IpAddress)
	if err != nil {
		return err
	}

	return nil
}

func Run(conf *conf.ConfFile, machine string, command string) ([]byte, error) {
	serv, err := conf.FindServerUnitByName(machine)
	if err != nil {
		return []byte{}, err
	}

	res, err := sshc.RunCommand(serv, command) // TODO: concatener les arguments
	if err != nil {
		return []byte{}, err
	}

	return res, nil
}
