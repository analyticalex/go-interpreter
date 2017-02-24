/*
Package main starts greets the user and starts repl.
*/

package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/analyticalex/go-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", 		user.Username)
	fmt.Printf("Please type commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}