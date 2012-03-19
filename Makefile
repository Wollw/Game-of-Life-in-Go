# $APPDIR/Makefile
TARG = main
all: clean
	$(MAKE) -C life
	$(MAKE) -f Makefile.$(TARG)
clean:
	$(MAKE) -C life clean
	$(MAKE) -f Makefile.$(TARG) clean
	rm -f $(TARG)
test:
	$(MAKE)
	./test.sh
	$(MAKE) clean 
