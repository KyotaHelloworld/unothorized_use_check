linux:
	go build -o sample.out main.go
mac:
	GOOS=darwin GOARCH=arm64  go build -o sample-mac.out main.go
	zip sample-mac.out.zip sample-mac.out