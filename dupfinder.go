package main

import (
	"fmt"; 
	"os";
	"flag"
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s root-directory\n", os.Args[0]);
	flag.PrintDefaults();
}

func main() {
	if flag.NArg() < 1 {
		Usage()
	}
	else {
		fmt.Printf("processing under %s\n", flag.Arg(0));
		fileMetadata := make(chan FileData);
		go ChecksumFiles(flag.Arg(0), fileMetadata);
		for fd := range fileMetadata {
			fmt.Printf("  found file %s with hash\n", fd.Name);
		}
	}
}

