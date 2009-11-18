include $(GOROOT)/src/Make.$(GOARCH)

TARG=dupfinder
GOFILES=\
	dupfinder.go\
	pathiterator.go\
	checksum.go

include $(GOROOT)/src/Make.cmd
