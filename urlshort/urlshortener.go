package main

import (
	"fmt"
	"net/http"
	//"time"
	//"math/rand"
)

func main() {
	fmt.Println("Hello world")
}

func get(link string) {
	resp, err := http.Get(link)

	if err != nil {
		fmt.Println("Error getting")
	}
	defer resp.Body.Close()
}
