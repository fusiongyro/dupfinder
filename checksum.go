package main

import (
	"crypto/sha1"; 
	"os"; 
	"io";
	"bytes";
	"encoding/base64"
)

const NCPU = 4

type Checksum struct {
	Name string;
	Hash string
}

type ChecksumIterator Path

// Gives us a handy way to satisfy Iterable and produce checksums
func (path ChecksumIterator) Iter() <-chan Checksum {	
	res := make(chan Checksum);
	
	// this largish function handles all of the fun in another goroutine
	go func() {
		// make our semaphore with the number of CPUs we want to support
		sem := make(chan int, NCPU);
		
		// prepopulate the semaphore with the thread IDs (cpu 1, cpu 2, etc.)
		for i := 0; i < NCPU; i += 1 { sem <- i }
		
		// now, iterate through the files we find and dispatch a process for
		// each one, using whatever CPU ID we happen to have handy
		for file := range Path(path).Iter() {
			go process(file, res, sem, <- sem);
		}
		
		// when we get here, we've run out of files to checksum, so we can close
		// the iteration
		close(res);		
	}();
	
	return res;
}

func process(path Path, out chan<- Checksum, res chan<- int, threadId int) {
	// ensure that we signal the semaphore once we're done
	defer func() { res <- threadId }();	
	
	// send the Checksum out the channel if there was no error
	if data, err := getHash(path); err == nil {	
		out <- Checksum{path.String(), data}	
	}
}

// Given a path string and a Hash interface, calculates the hash of the file's
// content. Uses Go's cool named result trick
func getHash(path Path) (result string, err os.Error) {
	// make a new hash calculator
	hash := sha1.New();
	
	// if we can open the file...
	if file, err := os.Open(string(path), os.O_RDONLY, 0); err == nil {
		// and if we can copy its contents to the hash
		if _, err := io.Copy(hash, file); err == nil {
			// then we have a result
			result = encodeBase64(hash.Sum())
		}
	}
	// return, whether we have nil for result or error
	return
}

// this is a fairly nasty way of getting a string out of a byte array in Base64
func encodeBase64(source []byte) string {
	// make a byte slice just big enough for the result of the encode operation
	dest := make([]byte, base64.StdEncoding.EncodedLen(len(source)));
	// encode it
	base64.StdEncoding.Encode(dest, source);
	// convert this byte buffer to a string
	return bytes.NewBuffer(dest).String();
}
