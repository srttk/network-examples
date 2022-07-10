package examples

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/saratonite/network-examples/utils"
)

func TcpClient(port string) {

	con, err := net.Dial("tcp", ":"+port)

	utils.CheckError(err, "Cant connect")

	defer con.Close()

	// Open os.Stdin and return a file handle

	inputReader := bufio.NewReader(os.Stdin)

	for {
		// Read user input

		fmt.Println("Enter the message to sent : ")
		// input wrap end

		input, _ := inputReader.ReadString('\n')

		inputInfo := strings.Trim(input, "\r\n")

		// If you enter q, exit

		if strings.ToUpper(inputInfo) == "Q" {
			fmt.Println("Stop input and disconnect")

			return
		}

		// Sent data

		fmt.Println("Start sending data")

		_, err = con.Write([]byte(inputInfo))

		if err != nil {
			return
		}
		fmt.Println("Sending suceeded")

		buf := [512]byte{}

		n, err := con.Read(buf[:])

		if err != nil {
			fmt.Println("Receice failed ", err)
			return
		}

		fmt.Println("Received as reply from the server :")

		fmt.Println(string(buf[:n]))
	}

}
