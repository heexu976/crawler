/*
 * @Date: 2023-10-09 21:49:40
 * @LastEditors: HeXu
 * @LastEditTime: 2023-10-10 21:15:47
 * @FilePath: /crawler/simpleHttpServer.go
 */
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "HELLO!")
}

// func main() {
// 	http.HandleFunc("/", hello)
// 	http.ListenAndServe("0.0.0.0:8080", nil)
// }
