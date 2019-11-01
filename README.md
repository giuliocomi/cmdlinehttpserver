# cmdlinehttpserver
Minimalistic Golang HttpServer

Simple standalone http server with directory listing enabled to allow a quick transfer of files.
It is ideal to compile and run it from 32/64 bits windows machines when all common transfer method (smb, ftp, tftp, .vbs, certutil, mshta, .., .ps1, nc, bitsadmin, etc.) are not an option.

Usage:
> .\cmdlinehttpserver.exe -a 0.0.0.0 -p 4343

To cross compile the standalone for different OS or architectures see https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04#step-4-%E2%80%94-building-executables-for-different-architectures

Example compiling windows executable from linux:

$ GOOS=windows GOARCH=386 go build -o cmdlinehttpserver.exe cmdlinehttpserver.go

By default it listens on port 8080 of all interfaces.

Security note: this commandline tool is everything but secure (http, directory listing, no auth, no rate limiting, etc.), for a more secure version with upload feature available have a look at https://github.com/giuliocomi/http-server.
