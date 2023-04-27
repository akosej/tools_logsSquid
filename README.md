# tools_logsSquid

Extract information from the access.log file

### Run app

`go run main.go -f access.log`

### Build for Windows

`GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -ldflags "-H windowsgui" -o tlsquid.exe ./main.go`
