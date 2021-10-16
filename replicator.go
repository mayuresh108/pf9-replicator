package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/platform9/pf9-replicator/grpc"
	"golang.org/x/crypto/ssh"
)

type ServerConnInfo struct {
	Server string
	Port   string
	User   string
	Key    string
}

func main() {

	// Read params from config file, config fiel will include: target host IP, username, ssh key path, du_fqdn, list of commands to be executed on host
	// 1. Extract embedded yaml, binary (client to copied, the client itself will have embedded yamls / binaries)
	// 2. Scp binary over to the host, start it on the host with ssh
	// 3. The client binary running on host will try to connect to this one, pass DU fqdn as argument to client
	// 4. open grpc server,
	//		- client invokes function SayHello on server
	//		- in response, return commands from the config file one by one
	//		- client executes the command and sends its output as arg in the next hello
	//		- last command 'exit' will result in client exiting, server too exits after this

	// Two use cases can be presented, updating a binary on the host (which was embedded in the client, extracted on host )
	// and updating addon operator image tag, lets say from 3.1.0 -> 4.0.0

	/*sci := ServerConnInfo{
		"10.128.240.149",
		"22",
		"centos",
		`id_rsa`,
	}*/

	/*clientConfig, _ := auth.PrivateKey("centos", "id_rsa", ssh.InsecureIgnoreHostKey())
	client := scp.NewClient("10.128.240.149:22", &clientConfig)
	err := client.Connect()
	if err != nil {
		fmt.Println("Couldn't establish a connection to the remote server ", err)
		return
	}
	defer client.Close()

	f, _ := os.Open("/tmp/hello.txt")
	defer f.Close()

	err = client.CopyFile(f, "/tmp/hello.txt", "0655")
	if err != nil {
		fmt.Println("Error while copying file ", err)
	}*/

	/*command := "whoami"
	log.Printf("Running command: %s", command)
	output, exitError := SSHCommandString(command, sci)
	fmt.Printf("Result: %s", output)
	fmt.Printf("Error: %s", exitError)*/

	go grpc.Server()
	for {
		grpc.Client()
	}
}

func (c *ServerConnInfo) Socket() string {
	return fmt.Sprintf("%s:%s", c.Server, c.Port)
}

func publicKeyFile(file string) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(key), nil
}

func generateSession(s ServerConnInfo) (*ssh.Session, ssh.Conn, error) {
	log.Print("Generating session...")
	log.Print("Reading public key")

	publicKey, err := publicKeyFile(s.Key)
	if err != nil {
		return nil, nil, err
	}

	config := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			publicKey,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	log.Print("Dialing ssh")

	conn, err := ssh.Dial("tcp", s.Socket(), config)
	if err != nil {
		return nil, nil, err
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := conn.NewSession()
	if err != nil {
		return nil, conn, err
	}

	log.Print("Returning ssh session")
	return session, conn, nil
}

func SSHCommandBool(command string, sci ServerConnInfo) (bool, error) {
	session, conn, err := generateSession(sci)
	if err != nil {
		if conn != nil {
			conn.Close()
		}

		return false, err
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	err = session.Run(command)

	session.Close()
	conn.Close()

	if err != nil {
		return false, err
	}
	return true, nil
}

func SSHCommandString(command string, sci ServerConnInfo) (string, error) {
	session, conn, err := generateSession(sci)
	if err != nil {
		if conn != nil {
			conn.Close()
		}

		return "", err
	}

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	log.Print("Running command on ssh")
	err = session.Run(command)

	session.Close()
	conn.Close()

	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(stdoutBuf.String(), "\n"), nil
}
