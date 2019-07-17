# cmdlinehttpserver
Minimalistic Golang HttpServer

Simple standalone http server with directory listing enabled to allow a quick transfer of files.
It is ideal to compile and run it from 32/64 bits windows machines when setting up a SMB share or using the FTP embedded command line client are not an option.

Usage:
> .\cmdlinehttpserver.exe -a 0.0.0.0 -p 4343

To cross compile the standalone for different OS or architectures:

$ GOOS=windows GOARCH=386 go build -o cmdlinehttpserver.exe cmdlinehttpserver.go

By default it listens on port 8080 of all interfaces.

Security note: this commandline tool is everything but secure (http, directory listing, no auth, etc.), for a more secure version with upload feature available have a look at https://github.com/giuliocomi/http-server.
