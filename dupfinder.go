package main

import "container/vector"

type DuplicateMap map[string] *vector.StringVector

func FindDuplicates(p Path) DuplicateMap {
	duplicates := make(map[string] *vector.StringVector);
	
	for fd := range ChecksumIterator(p).Iter() {
		vec, ok := duplicates[fd.Hash]; 
		if !ok {
			vec = vector.NewStringVector(0);
			duplicates[fd.Hash] = vec;
		}
		vec.Push(fd.Name);
	}
	return duplicates;
}
