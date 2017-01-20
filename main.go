package main

import (
	"fmt"

	"github.com/chazsmi/decorator-pattern/sshclient"

	"golang.org/x/crypto/ssh"
)

func main() {
	sshConfig := sshclient.Decorate(
		sshclient.User("charlie.smith1"),
	)

	connection, err := ssh.Dial("tcp", "127.0.0.1:22", sshConfig)
	if err != nil {
		fmt.Errorf("Failed to dial: %s", err)
	}
	fmt.Printf("%+v", connection)
}
