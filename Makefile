include $(GOROOT)/src/Make.$(GOARCH)

TARG=dupfinder
GOFILES=\
	dupfinder.go\
	dups.go\
	pathiterator.go\
	threadpool.go

include $(GOROOT)/src/Make.cmd
