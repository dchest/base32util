// Public domain.

// Encodes things read from stdin into base32.
package main

import (
	"encoding/base32"
	"fmt"
	"io"
	"os"
)

func main() {
	enc := base32.NewEncoder(base32.StdEncoding, os.Stdout)
	if _, err := io.Copy(enc, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	enc.Close()
	fmt.Println()
}
