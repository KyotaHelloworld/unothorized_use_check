linux:
	go build -o sample/alert.out alert/main.go
mac:
	GOOS=darwin GOARCH=arm64 go build -o sample/alert.out alert/main.go