build:
	@cp -r resources bin/resources
	@# Windows Build
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o bin/poketcgsim.exe main.go
	@# Linux Build
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o bin/poketcgsim main.go
