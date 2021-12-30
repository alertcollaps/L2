package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Create("test.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer resp.Body.Close()
	write := bufio.NewWriter(file)

	for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		write.WriteString(string(bs[:n]))

		if n == 0 || err != nil {
			break
		}
	}
}
