package main

import (
	"fmt"

	"github.com/Nuage-Malin/server-manager/hsmlib/conf"
	"github.com/Nuage-Malin/server-manager/hsmlib/sshc"
)

func run() error {
	config, err := conf.Load("mock/testconffile.json")
	if err != nil {
		return err
	}

	servConf, err := config.FindServerUnitByName("NuageCoquin")
	if err != nil {
		return err
	}

	res, err := sshc.RunCommand("systemctl suspend") // TODO: concatener les arguments
	if err != nil {
		return err
	}

	fmt.Print(string(res))

	return nil
}
