include $(GOROOT)/src/Make.$(GOARCH)

.SUFFIXES: .go .$O

OBJS=factorial.$O main.$O 
TARG=factorial

GOFILES=factorial.go

factorial: $(OBJS)
	$(LD) -o factorial $(OBJS)

clean:
	rm -f $(OBJS) factorial

.go.$O:
	$(GC) $<
	
test:
	gotest

testpackage-clean:
	rm -f _test/$(TARG).a _gotest_.$O

testpackage: _test/$(TARG).a

_test/$(TARG).a: _gotest_.$O $(OFILES)
	@mkdir -p _test/$(dir)
	rm -f _test/$(TARG).a
	gopack grc $@ _gotest_.$O $(OFILES)

_gotest_.$O: $(GOFILES) $(GOTESTFILES) $(PREREQ)
	$(GC) -o $@ $(GOFILES) $(GOTESTFILES)

importpath:
	@echo $(TARG)
