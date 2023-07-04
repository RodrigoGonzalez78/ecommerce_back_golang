package main

import "net/http"

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", Home())

}
