package main

import ("path"; "os")

// Not entirely sure that the type hell I wind up in elsewhere because of this
// is worth it.
type Path string

// satisfies exp/iterable.Iterable interface
func (p Path) Iter() <-chan Path {
	c := make(chan Path);
	go func() { 
		path.Walk(p.String(), pathIter(c), nil);
		close(c)
	}();
	return c
}

type pathIter chan<- Path

func (c pathIter) VisitDir(path string, d *os.Dir) bool {
	return path[0] != '.'
}

func (c pathIter) VisitFile(path string, d *os.Dir) {
	c <- Path(path)
}

func (p Path) String() string { return string(p); }