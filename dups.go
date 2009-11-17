package main

import ("path"; "hash"; "crypto/sha1"; "os"; "io")

type FileData struct {
	Name string;
	Hash []byte
}

type FileDataChannel chan FileData

// Returns a channel of FileDatas under a given path.
func ChecksumFiles(under string, out chan FileData) {
	path.Walk(under, FileDataChannel(out), nil);
	close(out);
}

// Methods for FileData channels to implement path.Visitor
func (FileDataChannel) VisitDir(path string, d *os.Dir) bool {
	return path[0] != '.'
}

func (c FileDataChannel) VisitFile(path string, d *os.Dir) {
	// calculate the hash
	hashAlgorithm := sha1.New();
	if hash, err := getHash(hashAlgorithm, path); err == nil {
		c <- FileData{path, hash};
	}
}

// Given a path string and a Hash interface, calculates the hash of the file's content
func getHash(hash hash.Hash, filename string) (result []byte, err os.Error) {
	if file, err := os.Open(filename, os.O_RDONLY, 0); err == nil {
		if _, err := io.Copy(hash, file); err == nil {
			result = hash.Sum()
		}
	}
	return
}

