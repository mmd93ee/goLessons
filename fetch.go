package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, urlin := range os.Args[1:] {
		fmt.Fprintf(os.Stdout, "Processing %s \n", urlin)

		if !strings.HasPrefix(urlin, "http://") {
			urlin = "http://" + urlin
		}

		resp, err := http.Get(urlin)
		if err != nil {
			errorHandler(urlin, err)
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errorHandler(urlin, err)
		}
		resp.Body.Close()

		r, err := ioutil.Readall(resp.StatusCode)
		if err != nil {
			errorHandler(urlin, err)
		}

		resp.Body.Close()

		fmt.Printf("%s", b)
	}
}

func errorHandler(u string, e error) {
	fmt.Fprintf(os.Stderr, "ERROR: Problem processing %s: %v", u, e)
	os.Exit(1)
}
