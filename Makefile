all: darwin linux windows

darwin:
	GOOS=darwin GOARC=amd64 go build -o build/darwin/gm ./main.go
linux:
	GOOS=linux GOARC=amd64 go build -o build/linux/gm ./main.go
windows:
	GOOS=windows GOARC=amd64 go build -o build/windows/gm.exe ./main.go
