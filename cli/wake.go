package main

import (
	"github.com/urfave/cli/v2"
	"nuagemalin.com/hsm/hsmlib/conf"
	"nuagemalin.com/hsm/hsmlib/wol"
)

func wake(cCtx *cli.Context) error {
	config, err := conf.Load(cCtx.String("file"))
	if err != nil {
		return err
	}

	servConf, err := config.FindServerUnitByName(cCtx.Args().First())
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
