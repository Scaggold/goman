all: darwin linux windows

darwin:
	GOPATH=`pwd` GOOS=darwin GOARC=amd64 go build -o build/darwin/gm ./main.go
linux:
	GOPATH=`pwd` GOOS=linux GOARC=amd64 go build -o build/linux/gm ./main.go
windows:
	GOPATH=`pwd` GOOS=windows GOARC=amd64 go build -o build/windows/gm.exe ./main.go
