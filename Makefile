include $(GOROOT)/src/Make.inc

TARG=dupfinder
GOFILES=\
	dupfinder.go\
	pathiterator.go\
	checksum.go\
	main.go

include $(GOROOT)/src/Make.cmd
