package main

import (
	"fmt"
	"os"
)

func main() {
	a := newApp()
	err := a.run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error running the app: ", err)
	}
}
