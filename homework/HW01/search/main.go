package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://62.152.59.9")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

}
