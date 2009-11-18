package main

import (
	"fmt"; 
	"os";
	"flag";
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
		// this is the cute part where I use it
		for fd := range ChecksumIterator(flag.Arg(0)).Iter() {
			fmt.Printf("  %s:  %s\n", fd.Name, fd.Hash);
		}
	}
}
