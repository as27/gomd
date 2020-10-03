package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	a := newApp()
	w, ok := a.bottom.(io.Writer)
	if ok {
		log.SetOutput(w)
	}
	err := a.run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error running the app: ", err)
	}
}
