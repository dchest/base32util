// Public domain.

// Encodes things read from stdin into base32.
package main

import (
	"bytes"
	"encoding/base32"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	fLowerCase   = flag.Bool("l", false, "lower case")
	fTrimPadding = flag.Bool("p", false, "remove padding")
)

func main() {
	flag.Parse()
	var b bytes.Buffer
	enc := base32.NewEncoder(base32.StdEncoding, &b)
	if _, err := io.Copy(enc, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	enc.Close()
	s := b.String()
	if *fLowerCase {
		s = strings.ToLower(s)
	}
	if *fTrimPadding {
		s = strings.TrimRight(s, "=")
	}
	fmt.Println(s)
}
