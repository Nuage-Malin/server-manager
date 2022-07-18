package conf

import (
	"encoding/json"
	"errors"
	"os"
)

type ServerUnit struct {
	Name       string `json:"name"`
	SshKey     string `json:"ssh_key"`
	IpAddress  string `json:"ip_address"`
	MacAddress string `json:"mac_address"`
}

type ConfFile struct {
	Servers []ServerUnit `json:"servers"`
}

func (c *ConfFile) FindServerUnitByName(name string) (*ServerUnit, error) {
	for _, server := range c.Servers {
		if server.Name == name {
			return &server, nil
		}
	}
	return nil, errors.New("Could not find server by name.")
}

// Load a configuration file located at the given filePath
func Load(filePath string) (ConfFile, error) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return ConfFile{}, err
	}

	var confFile ConfFile
	err = json.Unmarshal(dat, &confFile)
	if err != nil {
		return ConfFile{}, err
	}

	return confFile, nil
}
