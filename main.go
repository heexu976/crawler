package main

import (
	"fmt"
	"github.com/heexu976/crawler/collect"
)

func main() {
	url := "https://studygolang.com/"
	f := collect.BrowserFetch{}
	body, err := f.Get(url)
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	fmt.Println(string(body))
}
