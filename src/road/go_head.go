package main

import (
	"net/http"
	"io"
	"os"
)

func main() {
	resp, err := http.Head("http://www.aliyoyo.com/")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
