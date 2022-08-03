package sshc

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
	"nuagemalin.com/hsm/hsm/conf"
)

// Make a ssh.ClientConfig struct from a conf.ServerUnit
func MakeConfig(server *conf.ServerUnit) (*ssh.ClientConfig, error) {
	signer, err := ssh.ParsePrivateKey([]byte(server.SshKey))
	if err != nil {
		return nil, err
	}
	clientConf := &ssh.ClientConfig{
		User: server.SshUsername,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}
	return clientConf, nil
}

// Make a ssh.Client struct from a conf.ServerUnit
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

// Make a ssh.Session struct from a conf.ServerUnit
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
