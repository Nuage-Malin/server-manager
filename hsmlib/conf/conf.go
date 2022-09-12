package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type ServerUnit struct {
	Name        string  `json:"name"`
	SshUsername string  `json:"ssh_username"`
	SshKey      *string `json:"ssh_key,omitempty"`
	SshPassword *string `json:"ssh_password,omitempty"`
	IpAddress   string  `json:"ip_address"`
	SshPort     int     `json:"ssh_port"`
	MacAddress  string  `json:"mac_address"`
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
		if serv.SshKey == nil && serv.SshPassword == nil {
			panic(fmt.Sprintf("Server %s have empty password and key.", serv.Name))
		}
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

	SanityCheck(&confFile)

	return confFile, nil
}
