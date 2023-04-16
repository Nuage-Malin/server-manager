package main

import (
	"github.com/Nuage-Malin/server-manager/hsmlib/conf"
	"github.com/Nuage-Malin/server-manager/hsmlib/wol"
	"github.com/urfave/cli/v2"
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
