// Public domain.

// Encodes things read from stdin into base32.
package main

import (
	"encoding/base32"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Println(base32.StdEncoding.EncodeToString(b))

}
