package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/user"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type ConfigurationFile struct {
	Certificates []struct {
		Name  string   `json:"name"`
		Files []string `json:"files"`
	} `json:"certificates"`
	Tunnels []struct {
		Name        string `json:"name"`
		Bastion     string `json:"bastion"`
		Address     string `json:"address"`
		Localport   string `json:"localPort"`
		Certificate string `json:"certificate"`
	} `json:"tunnels"`
}

type Tunnel struct {
	Bastion          string   `json:"bastion"`
	Address          string   `json:"address"`
	Localport        string   `json:"localPort"`
	CertificatesPath []string `json:"-"`
}

type TunnelsMap map[string]*Tunnel

/** Idea from: https://ixday.github.io/post/golang_ssh_tunneling/ **/

const EnvSSHAuthSock = "SSH_AUTH_SOCK"

func authAgent() (ssh.AuthMethod, error) {
	conn, err := net.Dial("unix", os.Getenv(EnvSSHAuthSock))
	if err != nil {
		return nil, err
	}
	client, err := agent.NewClient(conn), err
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeysCallback(client.Signers), nil
}

func startTunnel(conn *ssh.Client, local, remote string) error {
	pipe := func(writer, reader net.Conn) {
		defer writer.Close()
		defer reader.Close()

		_, err := io.Copy(writer, reader)
		if err != nil {
			log.Printf("failed to copy: %s", err)
		}
	}
	listener, err := net.Listen("tcp", local)
	if err != nil {
		return err
	}
	for {
		here, err := listener.Accept()
		if err != nil {
			return err
		}
		go func(here net.Conn) {
			there, err := conn.Dial("tcp", remote)
			if err != nil {
				log.Fatalf("failed to dial to remote: %q", err)
			}
			go pipe(there, here)
			go pipe(here, there)
		}(here)
	}
}

func CreateTunnel() {
	// initiate auths methods
	authAgent, err := authAgent()
	if err != nil {
		log.Fatalf("failed to connect to the ssh agent: %q", err)
	}

	// here I am retrieving user from current execution,
	// you can pass it as argument if you want
	current, err := user.Current()
	if err != nil {
		log.Fatalf("failed to create ssh config: %q", err)
	}
	log.Printf("user.Current(): %q", current)

	// initialize SSH connection
	clientConfig := &ssh.ClientConfig{
		//User: current.Username,
		User: "<username>",
		Auth: []ssh.AuthMethod{authAgent},
		// you should not pass this option, but for the sake of simplicity
		// we use it here
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	clientConn, err := ssh.Dial("tcp", "<bastion host>:22", clientConfig)
	if err != nil {
		log.Fatalf("failed to connect to the ssh server: %q", err)
	}

	// tunnel traffic between local port 1600 and remote port 1500
	if err := startTunnel(clientConn, "localhost:8080", "<remote address>:8095"); err != nil {
		log.Fatalf("failed to tunnel traffic: %q", err)
	}
}
