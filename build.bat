@echo off
set glc_version=1.1.0

set GOOS=darwin
set GOARCH=amd64
go build -o ./output/glc_%glc_version%_darwin_amd64/glc

set GOOS=windows
set GOARCH=amd64
go build  -o ./output/glc_%glc_version%_windows_amd64/glc.exe

set GOOS=linux
set GOARCH=amd64
go build  -o ./output/glc_%glc_version%_linux_amd64/glc
