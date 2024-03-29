package sshc

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Nuage-Malin/server-manager/hsmlib/conf"
	"golang.org/x/crypto/ssh"
)

// Makes a ssh.ClientConfig struct from a conf.ServerUnit
func MakeConfig(server *conf.ServerUnit) (*ssh.ClientConfig, error) {
	var signer ssh.Signer

	clientConf := &ssh.ClientConfig{
		User:            server.SshUsername,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // TODO: Modify configuration to handle host key callback security
	}

	if server.SshKey != nil {
		file, err := ioutil.ReadFile(*server.SshKey)
		if err != nil {
			return nil, err
		}

		signer, err = ssh.ParsePrivateKey(file)
		if err != nil {
			return nil, err
		}
		clientConf.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		}
	} else if server.SshPassword != nil {
		clientConf.Auth = []ssh.AuthMethod{
			ssh.Password(*server.SshPassword),
		}
	}

	return clientConf, nil
}

// Makes a ssh.Client struct from a conf.ServerUnit
func ConnectToServer(server *conf.ServerUnit) (*ssh.Client, error) {
	conf, err := MakeConfig(server)
	if err != nil {
		return nil, err
	}

	client, err := ssh.Dial("tcp", strings.Join([]string{server.IpAddress, ":", fmt.Sprint(server.SshPort)}, ""), conf)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Makes a ssh.Session struct from a conf.ServerUnit
func MakeSession(server *conf.ServerUnit) (*ssh.Session, error) {
	client, err := ConnectToServer(server)
	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

// Runs command on ssh server and returns the output
func RunCommand(server *conf.ServerUnit, cmd string) ([]byte, error) {
	session, err := MakeSession(server)
	if err != nil {
		return nil, err
	}

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return nil, err
	}

	return output, nil
}
