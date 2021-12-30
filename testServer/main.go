package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type resultAndError struct {
	Result string `json:"result,omitempty"`
	Err    string `json:"error,omitempty"`
}

func main() {
	r := new(resultAndError)
	r.Result = "Falder"
	fmt.Println()

	json.NewDecoder()
	con, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatal(err)
	}
	accept, err := con.Accept()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	bufReader := bufio.NewReader(accept)
	for {
		rb, err := bufReader.ReadByte()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(rb))
	}
}
