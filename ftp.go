package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {

	DEFAULTDIR := "/"
	DEFAULTPORT := "2121"

	flag.Parse()

	if (len(flag.Args())) < 1 {
		fmt.Println("EXAMPLE:  ftp  /data  2121")
		os.Exit(0)
	}

	DEFAULTDIR = flag.Arg(0)
	DEFAULTPORT = flag.Arg(1)

	fmt.Println("DIR:", DEFAULTDIR, "is sharing ,Now Listening Local port:"+DEFAULTPORT)

	h := http.FileServer(http.Dir(DEFAULTDIR))
	http.ListenAndServe(":"+DEFAULTPORT, h)

}
