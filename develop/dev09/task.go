package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	url := flag.Arg(0)

	err := wget(url)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func wget(url string) error {

	if url == "" {
		return errors.New("empty url")
	}

	resp, err := http.Get(url)
	if err != nil {
		return errors.New("can't execute get request")
	}

	defer resp.Body.Close()

	file, err := os.Create("out.html")
	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	fmt.Println("Success")
	return nil
}
