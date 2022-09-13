package main

import (
	"flag"
	"fmt"

	x "github.com/saratonite/network-examples/examples"
)

func main() {

	example := flag.String("example", "none", "Name of the example")
	hostname := flag.String("hostname", "google.com", "Hostname ")
	port := flag.String("port", "5252", "Port")
	url := flag.String("url", "https://google.com", "URL")
	key := flag.String("key", "key.pem", "Public key file")

	flag.Parse()

	switch *example {

	case "tcp-connect":
		x.TcpConnect(*hostname)

	case "tcp-server":
		x.TcpServer(*port)

	case "tcp-client":
		x.TcpClient(*port)

	case "http-server":
		x.HttpServer(*port)

	case "http-client":
		x.HttpClient(*url)
	case "ssh-client":
		x.SshClient(*hostname, *key)
	default:
		fmt.Println("No example selected \nplease specify the exaple usig --example <example-name>")

	}

}
