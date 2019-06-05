package seed

import (
	"fmt"
	"github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func (app *App) Deploy() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	//Use SSH keys if possible.
	SSHPrivateKey, err := ioutil.ReadFile(usr.HomeDir + "/.ssh/id_rsa")
	if err != nil {
		return err
	}

	//Parse private key:
	signer, err := ssh.ParsePrivateKey(SSHPrivateKey)

	var Host = app.host + ":22"
	var ClientConfig = &ssh.ClientConfig{
		User: "root",

		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
			ssh.PasswordCallback(func() (string, error) {
				fmt.Print("Enter Password: \r")
				bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
				password := string(bytePassword)
				fmt.Println()
				return password, err
			}),
		},

		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	var client = scp.NewClient(Host, ClientConfig)

	err = client.Connect()
	if err != nil {
		return err
	}

	//Compile for linux.
	var cmd = exec.Command("go", "build", "-o", "server.build")
	cmd.Dir = Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(),
		"GOOS=linux",
	)

	err = cmd.Run()
	if err != nil {
		return err
	}

	//Send to the server.
	executable, err := os.Open("server.build")
	if err != nil {
		return err
	}

	//Run a few commands!
	{
		var client, err = ssh.Dial("tcp", Host, ClientConfig)
		if err != nil {
			return err
		}

		session, err := client.NewSession()
		if err != nil {
			return err
		}

		err = session.Run("rm /srv/https/" + app.host + "/" + app.host)
		if err != nil {
			return err
		}
		client.Close()
	}

	err = client.CopyFromFile(*executable, "/srv/https/"+app.host+"/"+app.host, "0655")
	if err != nil {
		return err
	}

	client.Close()
	executable.Close()

	//Run a few commands!
	{
		var client, err = ssh.Dial("tcp", Host, ClientConfig)
		if err != nil {
			return err
		}

		session, err := client.NewSession()
		if err != nil {
			return err
		}

		err = session.Run("systemctl restart " + app.host)
		if err != nil {
			return err
		}
		client.Close()
	}

	return nil
}
