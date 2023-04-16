package main

import (
	"fmt"

	"github.com/Nuage-Malin/server-manager/hsmlib/conf"
	"github.com/Nuage-Malin/server-manager/hsmlib/sshc"
	"github.com/urfave/cli/v2"
)

func run(cCtx *cli.Context) error {
	config, err := conf.Load(cCtx.String("file"))
	if err != nil {
		return err
	}

	servConf, err := config.FindServerUnitByName(cCtx.Args().Get(0))
	if err != nil {
		return err
	}

	res, err := sshc.RunCommand(servConf, cCtx.Args().Get(1)) // TODO: concatener les arguments
	if err != nil {
		return err
	}

	fmt.Print(string(res))

	return nil
}
