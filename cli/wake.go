package main

import (
	"github.com/Nuage-Malin/server-manager/hsmlib/conf"
	"github.com/Nuage-Malin/server-manager/hsmlib/wol"
)

func wake() error {
	config, err := conf.Load("mock/testconffile.json")
	if err != nil {
		return err
	}

	servConf, err := config.FindServerUnitByName("NuageCoquin")
	if err != nil {
		return err
	}

	magicPacket, err := wol.NewPacketFromString(servConf.MacAddress)
	if err != nil {
		return err
	}

	err = wol.SendPacket(magicPacket, servConf.IpAddress)
	if err != nil {
		return err
	}

	return nil
}
