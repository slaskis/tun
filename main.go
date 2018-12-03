package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net"
)

func main() {
	opts := struct {
		Listen             string
		Connect            string
		ServerName         string
		InsecureSkipVerify bool
	}{}
	flag.StringVar(&opts.Listen, "listen", "127.0.0.1:6060", "listening address")
	flag.StringVar(&opts.Connect, "connect", "", "connect address")
	flag.StringVar(&opts.ServerName, "servername", "", "remote tls servername")
	flag.BoolVar(&opts.InsecureSkipVerify, "insecure", false, "skip tls verify")
	flag.Parse()

	if opts.Connect == "" {
		flag.Usage()
		return
	}

	listener, err := net.Listen("tcp4", opts.Listen)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Printf("listen: %s\n", opts.Listen)

	for {
		rw, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("connect: %s\n", opts.Listen)

		go func() {
			defer rw.Close()

			cn, err := tls.Dial("tcp", opts.Connect, &tls.Config{
				ServerName:         opts.ServerName,
				InsecureSkipVerify: opts.InsecureSkipVerify,
			})
			if err != nil {
				log.Fatal(err)
			}

			if err := Pipe(cn, rw); err != nil {
				log.Fatal(err)
			}
		}()
	}
}
