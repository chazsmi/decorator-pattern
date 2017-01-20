package sshclient

import "golang.org/x/crypto/ssh"

type Client interface {
	Do(*ssh.ClientConfig) (*ssh.Client, error)
}

type ClientFunc func(*ssh.ClientConfig) (*ssh.Client, error)

func (f ClientFunc) Do(c *ssh.ClientConfig) (*ssh.Client, error) {
	return f(c)
}
