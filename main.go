package main

import (
	"flag";
	"fmt";
	"os"
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
	} else {
		// this is the cute part where I use it
		duplicates := FindDuplicates(Path(flag.Arg(0)));
		// now walk through all my stuff and print out files which are duplicates and
		// what they're duplicated with
		var i int;
		for hash, filenames := range duplicates {
			if filenames.Len() > 1 {
				fmt.Printf("Duplicate files with hash %s:\n", hash);
				i = 1;
				for filename := range filenames.Iter() {
					fmt.Printf("  %d. %s\n", i, filename);
					i++
				}
				println()
			}
		}
	}
}
