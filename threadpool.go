package main

func process(path Path, out chan FileData, res chan int, threadId int) {
	// ensure that we signal the semaphore once we're done
	defer func() { res <- threadId }();
	
	// try to get a checksum for this file
	data, err := getHash(path);
	
	// send the filedata out the channel if there was no error
	if err == nil {	out <- FileData{path.String(), data}	}
}

const NCPU = 4

func FileDataIterator(path Path) chan FileData {	
	res := make(chan FileData);
	
	var threadId int;
	
	// this largish function handles all of the fun in another goroutine
	go func() {
		// make our semaphore with the number of CPUs we want to support
		sem := make(chan int, NCPU);
		
		// prepopulate the semaphore with the thread IDs (cpu 1, cpu 2, etc.)
		for i := 0; i < NCPU; i += 1 { sem <- i }
		
		// now, iterate through the files we find and dispatch a process for
		// each one, using whatever CPU ID we happen to have handy
		for file := range path.Iter() {
			threadId = <- sem;
			go process(file, res, sem, threadId);
		}
		
		// when we get here, we've run out of files to checksum, so we can close
		// the iteration
		close(res);		
	}();
	
	return res;
}