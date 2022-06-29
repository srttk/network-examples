package examples

import (
	"fmt"
	"log"
	"net"
)

func TcpConnect(hostname string) {
	con, error := net.Dial("tcp", hostname+":http")

	if error != nil {

		log.Fatalln("Cant connect", error)
	}
	defer con.Close()

	fmt.Println("Connectecd to ", hostname)

}
