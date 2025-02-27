package main

import (
	"01/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands, type 'EXIT' to quit\n")
	repl.Start(os.Stdin, os.Stdout)
}
