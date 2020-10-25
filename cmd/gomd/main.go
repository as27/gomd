package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	flagVerbose = flag.Bool("verbose", false, "start in verbose mode. Will print internal logs in the app. Use this for debugging purposes.")
)

func main() {
	flag.Parse()
	a := newApp()
	w, ok := a.log.(io.Writer)
	if ok {
		if *flagVerbose {
			log.SetOutput(w)
		}
	}
	if !*flagVerbose {
		log.SetOutput(ioutil.Discard)
	}
	a.conf.verbose = *flagVerbose
	err := a.run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error running the app: ", err)
	}
}
