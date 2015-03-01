// Public domain.

// Encodes things read from stdin into base32.
package main

import (
	"encoding/base32"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	fLowerCase   = flag.Bool("l", false, "lower case")
	fTrimPadding = flag.Bool("p", false, "remove padding")
)

func main() {
	flag.Parse()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	s := base32.StdEncoding.EncodeToString(b)
	if *fLowerCase {
		s = strings.ToLower(s)
	}
	if *fTrimPadding {
		s = strings.TrimRight(s, "=")
	}
	fmt.Println(s)
}
