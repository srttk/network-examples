package examples

import (
	"fmt"
	"io"
	"log"

	"github.com/gliderlabs/ssh"
)

func SshServer() {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello %v", s.User()))
	})

	log.Print("\nSSH Server running on port 22")

	log.Fatal(ssh.ListenAndServe(":22", nil))
}
