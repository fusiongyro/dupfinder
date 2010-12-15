package main

import ("path"; "os")

type Path string

// satisfies exp/iterable.Iterable interface
func (p Path) Iter() <-chan Path {
	// this is our result channel
	c := make(chan Path)
	
	// this goroutine recursively walks the path and then closes this channel
	go func() { 
		path.Walk(p.String(), pathIter(c), nil)
		close(c)
	}();
	
	// return our result channel
	return c
}

type pathIter chan<- Path

// path.Visitor interface definitions
func (c pathIter) VisitDir(path string, d *os.FileInfo) bool {
	return path[0] != '.'
}

func (c pathIter) VisitFile(path string, d *os.FileInfo) {
	c <- Path(path)
}

// for the sake of convenience and not abusing type casts
func (p Path) String() string { return string(p) }
