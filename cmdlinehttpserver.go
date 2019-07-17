package main

import (
"bytes"
"flag"
"fmt"
"net/http"
)

func main() {
var port = flag.String("p", "8080", "The port to run on")
var address = flag.String("a", "0.0.0.0", "The interface to listen on")
flag.Parse()

schema := bytes.Buffer{}
schema.WriteString(*address)
schema.WriteString(":")
schema.WriteString(*port)

err := DirList(".", "/")
if err != nil {
	fmt.Print(err.Error())
}

http.ListenAndServe(schema.String(), nil)
}

func DirList(localfolder, webpath string) (err error) {
fs := http.FileServer(http.Dir(localfolder))
http.Handle(webpath, http.StripPrefix(webpath, fs))
return
}