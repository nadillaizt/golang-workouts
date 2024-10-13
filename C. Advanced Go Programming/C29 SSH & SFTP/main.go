package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Function to read and return an SSH authentication method using a private key file
func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := os.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}

func main() {
	// Example 1: Connect via SSH and execute a single command
	func() {
		const SSH_ADDRESS = "0.0.0.0:22"
		const SSH_USERNAME = "user"
		const SSH_PASSWORD = "password"

		// SSH client configuration
		sshConfig := &ssh.ClientConfig{
			User:            SSH_USERNAME,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Auth: []ssh.AuthMethod{
				ssh.Password(SSH_PASSWORD),
				// PublicKeyFile(SSH_KEY),
			},
		}

		// Establish SSH connection
		client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
		if client != nil {
			defer client.Close()
		}
		if err != nil {
			log.Fatal("Failed to establish connection: " + err.Error())
		}

		// Start a new SSH session
		session, err := client.NewSession()
		if session != nil {
			defer session.Close()
		}
		if err != nil {
			log.Fatal("Failed to start session: " + err.Error())
		}

		// Set up I/O redirection
		session.Stdin = os.Stdin
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr

		// Execute a remote command
		err = session.Run("ls -l ~/")
		if err != nil {
			log.Fatal("Command execution failed: " + err.Error())
		}
	}()

	// Example 2: Send multiple commands over an SSH session
	func() {
		const SSH_ADDRESS = "0.0.0.0:22"
		const SSH_USERNAME = "user"
		const SSH_PASSWORD = "password"

		// SSH client configuration
		sshConfig := &ssh.ClientConfig{
			User:            SSH_USERNAME,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Auth: []ssh.AuthMethod{
				ssh.Password(SSH_PASSWORD),
				// PublicKeyFile(SSH_KEY),
			},
		}

		// Establish SSH connection
		client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
		if err != nil {
			log.Fatal("Failed to establish connection: " + err.Error())
		}

		// Start a new SSH session
		session, err := client.NewSession()
		if session != nil {
			defer session.Close()
		}
		if err != nil {
			log.Fatal("Failed to start session: " + err.Error())
		}

		var stdout, stderr bytes.Buffer
		session.Stdout = &stdout
		session.Stderr = &stderr

		// Create input stream for sending multiple commands
		stdin, err := session.StdinPipe()
		if err != nil {
			log.Fatal("Error getting input stream: " + err.Error())
		}

		// Start bash shell
		err = session.Start("/bin/bash")
		if err != nil {
			log.Fatal("Failed to start shell: " + err.Error())
		}

		// Send multiple commands to the session
		commands := []string{
			"cd /where/is/the/gopath",
			"cd src/myproject",
			"ls",
			"exit",
		}
		for _, cmd := range commands {
			if _, err = fmt.Fprintln(stdin, cmd); err != nil {
				log.Fatal(err)
			}
		}

		// Wait for commands to finish
		err = session.Wait()
		if err != nil {
			log.Fatal(err)
		}

		// Print output and error logs
		fmt.Println("============== ERROR")
		fmt.Println(strings.TrimSpace(stderr.String()))

		fmt.Println("============== OUTPUT")
		fmt.Println(strings.TrimSpace(stdout.String()))
	}()

	// Example 3: SFTP file upload via SSH
	func() {
		const SSH_ADDRESS = "go.eaciit.com:22"
		const SSH_USERNAME = "developer"
		const SSH_KEY = "/Users/novalagung/Documents/developer.pem"

		// SSH client configuration using private key authentication
		sshConfig := &ssh.ClientConfig{
			User:            SSH_USERNAME,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Auth: []ssh.AuthMethod{
				PublicKeyFile(SSH_KEY),
			},
		}

		// Establish SSH connection
		client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
		if client != nil {
			defer client.Close()
		}
		if err != nil {
			log.Fatal("Failed to establish connection: " + err.Error())
		}

		// Create SFTP client
		sftpClient, err := sftp.NewClient(client)
		if err != nil {
			log.Fatal("Failed to create SFTP client: " + err.Error())
		}

		// Open destination file on remote server
		fDestination, err := sftpClient.Create("/data/nginx/files/test-file.txt")
		if err != nil {
			log.Fatal("Failed to create destination file: " + err.Error())
		}

		// Open local source file
		fSource, err := os.Open("/Users/novalagung/Desktop/test-file.txt")
		if err != nil {
			log.Fatal("Failed to open source file: " + err.Error())
		}

		// Copy local file to remote server
		_, err = io.Copy(fDestination, fSource)
		if err != nil {
			log.Fatal("Failed to copy file: " + err.Error())
		}

		log.Println("File successfully copied.")
	}()
}
