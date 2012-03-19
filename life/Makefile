# $APPDIR/pkg1/Makefile
include $(GOROOT)/src/Make.inc
myall: all
	cp _obj/$(TARG).a ..
TARG        = life
GOFILES     = $(TARG).go
CLEANFILES  = ../$(TARG).a
include $(GOROOT)/src/Make.pkg 
