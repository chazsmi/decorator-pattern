package sshclient

import "golang.org/x/crypto/ssh"

func User(u string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(s *ssh.ClientConfig) (*ssh.Client, error) {
			s.User = u
			return c.Do(s)
		})
	}
}
