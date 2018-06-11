package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	//fmt.Println(resp)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	//lec 65
	//io.Copy(os.Stdout, resp.Body)

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	//return 1, nil
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes:", len(bs))
	return len(bs), nil
}
