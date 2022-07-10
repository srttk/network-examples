package examples

import (
	"fmt"
	"log"
	"net"
)

func TcpServer(port string) {

	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatal("Cant listen server ", err)

	}

	defer listener.Close()

	fmt.Println("TCP Server listeng on port ", port)

	for {

		connection, err := listener.Accept()

		if err != nil {
			log.Fatal("TCP Connection error ", err)

		}

		if connection != nil {
			fmt.Println("New incoming connection")
		}

		go process(connection)

	}
}

func process(con net.Conn) {

	defer con.Close()

	buf := make([]byte, 1024)

	len, err := con.Read(buf)

	if err != nil {
		fmt.Println("Buffer read error ")
		return
	}

	fmt.Println("Message received \n ", string(buf[:len]))

	con.Write([]byte("MESSAGE RECEIVED : OK"))

}
