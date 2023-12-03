package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main(){
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)
}