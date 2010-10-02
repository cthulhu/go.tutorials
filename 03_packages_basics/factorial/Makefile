include $(GOROOT)/src/Make.inc

TARG=factorial
GOFMT=gofmt -spaces=true -tabindent=false -tabwidth=4

GOFILES=\
  factorial.go\

include $(GOROOT)/src/Make.pkg

format:
	${GOFMT} -w factorial.go
	${GOFMT} -w main.go

