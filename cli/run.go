package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"nuagemalin.com/hsm/hsmlib/conf"
	"nuagemalin.com/hsm/hsmlib/sshc"
)

func run(cCtx *cli.Context) error {
	config, err := conf.Load(cCtx.String("file"))
	if err != nil {
		return err
	}

	servConf, err := config.FindServerUnitByName(cCtx.String("target"))
	if err != nil {
		return err
	}

	res, err := sshc.RunCommand(servConf, cCtx.Args().Get(0))
	if err != nil {
		return err
	}

	fmt.Print(string(res))

	return nil
}
