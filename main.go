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
	fAlphabet    = flag.String("a", "", `alphabet ("" = standard, "hex", "zooko", or alphabet characters)`)
	fGroup       = flag.Int("g", 0, "split into groups of N characters")
	fGroupSep    = flag.String("gs", " ", "group separator for -g option")
)

var encodings = map[string]*base32.Encoding{
	"":      base32.StdEncoding,
	"hex":   base32.HexEncoding,
	"zooko": base32.NewEncoding("YBNDRFG8EJKMCPQXOT1UWISZA345H769"),
}

func main() {
	flag.Parse()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	alpha := *fAlphabet
	enc, ok := encodings[alpha]
	if !ok {
		if len(alpha) != 32 {
			fmt.Fprintf(os.Stderr, "unknown alphabet: %s\n", alpha)
			os.Exit(2)
		}
		enc = base32.NewEncoding(alpha)
	}
	s := enc.EncodeToString(b)
	if *fLowerCase {
		s = strings.ToLower(s)
	}
	if *fTrimPadding {
		s = strings.TrimRight(s, "=")
	}
	g := *fGroup
	if g > 0 {
		rs := ""
		for i, r := range s {
			if i > 0 && i%g == 0 {
				rs += *fGroupSep
			}
			rs += string(r)
		}
		s = rs
	}
	fmt.Println(s)
}
