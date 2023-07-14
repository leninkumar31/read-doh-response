package main

import (
	"fmt"
	"log"
	"os"

	"github.com/miekg/dns"
)

func main() {

	argsWithProg := os.Args
	path := argsWithProg[1]

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening file", err)
	}

	defer file.Close()

	fmt.Printf("%s opened\n", path)
	data := make([]byte, 512)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	msg := &dns.Msg{}
	msg.Unpack(data)
	fmt.Println(msg.String())
}
