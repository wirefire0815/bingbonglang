package main

import (
	"bingbonglang/lexer"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi %s! This is bingbong bingbong bingbong\n", user.Username)

	lexer.Start(os.Stdin, os.Stdout)
}
