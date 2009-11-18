package main

import (
	"fmt"; 
	"os";
	"flag";
	"bytes";
	"encoding/base64"
)

// Display a user-friendly usage message
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s root-directory\n", os.Args[0]);
	flag.PrintDefaults();
}

// The main function. If we have enough arguments, use them.
func main() {
	if flag.NArg() < 1 {
		Usage()
	}
	else {
		path := Path(flag.Arg(0));
		for fd := range FileDataIterator(path) {
			fmt.Printf("  %s:  %s\n", fd.Name, encodeBase64(fd.Hash));
		}
	}
}

// this is a fairly nasty way of getting a string out of a byte array in Base64
func encodeBase64(source []byte) string {
	dest := make([]byte, base64.StdEncoding.EncodedLen(len(source)));
	base64.StdEncoding.Encode(dest, source);
	return bytes.NewBuffer(dest).String();
}