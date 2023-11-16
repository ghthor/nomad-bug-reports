.PHONY: all nobug bug

all: nobug

bug:
	sed -i '/replace/d' go.mod
	go run main.go

nobug: bug
	go run main.go
