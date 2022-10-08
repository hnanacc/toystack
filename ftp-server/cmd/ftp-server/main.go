package main

import (
	"flag"

	"github.com/bitbeast18/from-scratch/ftp-server/internal/server"
)

func main() {
	host := flag.String("host", "", "host address ipv4/ipv6")
	port := flag.String("port", "", "port to listen on")
	
	flag.Parse()

	server.New(*host, *port).Serve()
}