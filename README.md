# tools_logsSquid

Extract information from the access.log file

### General info

* You must copy the file you want to analyze to the application folder
* **You do not lose the original file.** The app generates a copy log.copy

### Build for Windows

`GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -ldflags "-H windowsgui" -o tlsquid.exe ./main.go`

### How to use the app in linux

```go
go build -o tlsquid main.go

./tlsquid -f name.log

//If you do not specify -f the application defaults to access.log

```

### Extract from the original log

```
1678045741.710  49875 10.26.90.55 TCP_MISS_ABORTED/000 0 POST http://c.whatsapp.net/chat ej ORIGINAL_DST/157.240.6.54 -
1678045779.610  99898 10.26.90.55 TCP_TUNNEL/200 5459 CONNECT 172.217.28.99:443 ej ORIGINAL_DST/172.217.28.99 -
1678045779.611 102969 10.26.90.55 TCP_TUNNEL/200 975966 CONNECT 172.217.173.36:443 ej ORIGINAL_DST/172.217.173.36 -
1678045779.612  99893 10.26.90.55 TCP_TUNNEL/200 5459 CONNECT 172.217.28.99:443 ej ORIGINAL_DST/172.217.28.99 -
1678045779.612  99362 10.26.90.55 TCP_TUNNEL/200 5458 CONNECT 172.217.28.99:443 ej ORIGINAL_DST/172.217.28.99 -
```

### Result

```
1678045741.710  49875 10.26.90.55 TCP_MISS_ABORTED/000 0 POST http://c.whatsapp.net/chat ej ORIGINAL_DST/whatsapp-chatd-edge-shv-01-bog1.facebook.com. -
1678045779.610  99898 10.26.90.55 TCP_TUNNEL/200 5459 CONNECT bog02s07-in-f3.1e100.net.:443 ej ORIGINAL_DST/bog02s07-in-f3.1e100.net. -
1678045779.611 102969 10.26.90.55 TCP_TUNNEL/200 975966 CONNECT bog02s12-in-f4.1e100.net.:443 ej ORIGINAL_DST/bog02s12-in-f4.1e100.net. -
1678045779.612  99893 10.26.90.55 TCP_TUNNEL/200 5459 CONNECT bog02s07-in-f3.1e100.net.:443 ej ORIGINAL_DST/bog02s07-in-f3.1e100.net. -
1678045779.612  99362 10.26.90.55 TCP_TUNNEL/200 5458 CONNECT bog02s07-in-f3.1e100.net.:443 ej ORIGINAL_DST/bog02s07-in-f3.1e100.net. -
1678045781.880 101327 10.26.90.55 TCP_TUNNEL/200 6952 CONNECT bog02s17-in-f3.1e100.net.:443 ej ORIGINAL_DST/bog02s17-in-f3.1e100.net. -
```
