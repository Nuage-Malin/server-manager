package conf

import (
	"encoding/json"
	"errors"
	"os"
)

type ServerUnit struct {
	Name        string `json:"name"`
	SshUsername string `json:"ssh_username"`
	SshKey      string `json:"ssh_key"`
	IpAddress   string `json:"ip_address"`
	SshPort     int    `json:"ssh_port"`
	MacAddress  string `json:"mac_address"`
}

type ConfFile struct {
	Servers []ServerUnit `json:"servers"`
}

func contains(s []string, f string) bool {
	for _, elem := range s {
		if elem == f {
			return true
		}
	}
	return false
}

// Check the sanity of the ConfFile, panic if an error is present
func SanityCheck(conf *ConfFile) {
	checked := make([]string, 0, len(conf.Servers))

	for _, serv := range conf.Servers {
		if contains(checked, serv.Name) {
			panic("2 servers have the same name.")
		}
		checked = append(checked, serv.Name)
	}
}

// Find Server by name.
// Returns the corresponding ServerUnit
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
