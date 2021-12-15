package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func getTime() (time.Time, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return time, err
}

func main() {
	time, _ := getTime()
	fmt.Printf("Time now is %v", time.Format("15:04:05"))
}
