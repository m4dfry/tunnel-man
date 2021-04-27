package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"

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
type TunnelsManager struct {
	buffer     chan string
	tunnelsMap TunnelsMap
}

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

// https://ixday.github.io/post/golang_ssh_tunneling/
func tunnel(conn *ssh.Client, local, remote string) error {
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

func (tm *TunnelsManager) CreateTunnel(name string) {
	t := tm.tunnelsMap[name]
	if t == nil {
		log.Fatalf("tunnel %s not found", name)
	}

	bsplit := strings.Split(t.Bastion, "@")
	if len(bsplit) < 2 {
		log.Fatalf("failed to read bastion data: %s", t.Bastion)
	}
	user := bsplit[0]
	bastionaddr := bsplit[1]

	// initiate auths methods
	authAgent, err := authAgent()
	if err != nil {
		log.Fatalf("failed to connect to the ssh agent: %q", err)
	}

	// initialize SSH connection
	clientConfig := &ssh.ClientConfig{
		//User: current.Username,
		User: user,
		Auth: []ssh.AuthMethod{authAgent},
		// you should not pass this option, but for the sake of simplicity
		// we use it here
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	clientConn, err := ssh.Dial("tcp", bastionaddr, clientConfig)
	if err != nil {
		log.Fatalf("failed to connect to the ssh server: %q", err)
	}

	// tunnel traffic between local port 1600 and remote port 1500
	if err := tunnel(clientConn, "localhost:"+t.Localport, t.Address); err != nil {
		log.Fatalf("failed to tunnel traffic: %q", err)
	}
}
