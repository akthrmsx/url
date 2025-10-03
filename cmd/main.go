package main

import (
	"fmt"
	"log"

	"github.com/akthrmsx/url"
)

func main() {
	u, err := url.Parse("https://github.com/akthrmsx")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("Host  :", u.Host)
	fmt.Println("Path  :", u.Path)
	fmt.Println(u)
}
