package main

import (
	"fmt"
	"os"
)

func main() {
	home := os.Getenv("HOME")
	user := os.Getenv("USER")
	path := os.Getenv("PATH")
	fmt.Printf("home=%s, user=%s path=% \n", home, user, path)
}
