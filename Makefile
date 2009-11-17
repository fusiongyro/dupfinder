include $(GOROOT)/src/Make.$(GOARCH)

TARG=dupfinder
GOFILES=\
	dupfinder.go\
	dups.go

include $(GOROOT)/src/Make.cmd
