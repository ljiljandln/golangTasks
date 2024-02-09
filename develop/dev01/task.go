package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func PrintTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	} else {
		fmt.Print(time)
	}
}

func main() {
	PrintTime()
}
