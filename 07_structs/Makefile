include $(GOROOT)/src/Make.$(GOARCH)
 
.SUFFIXES: .go .$O

OBJS=super_point.$O structs.$O 
TARG=structs

structs: $(OBJS)
	$(LD) -o structs $(OBJS)

clean:
	rm -f $(OBJS) structs

.go.$O:
	$(GC) $<

