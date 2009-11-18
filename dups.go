package main

import ("crypto/sha1"; "os"; "io")

type FileData struct {
	Name string;
	Hash []byte
}

// Given a path string and a Hash interface, calculates the hash of the file's content
func getHash(path Path) (result []byte, err os.Error) {
	hash := sha1.New();
	if file, err := os.Open(string(path), os.O_RDONLY, 0); err == nil {
		if _, err := io.Copy(hash, file); err == nil {
			result = hash.Sum()
		}
	}
	return
}

