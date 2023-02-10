//go:build !wasm
// +build !wasm

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/vugu/vugu/simplehttp"
)

func main() {
	dev := flag.Bool("dev", false, "Enable development features")
	dir := flag.String("dir", ".", "Project directory")
	httpl := flag.String("http", "127.0.0.1:8877", "Listen for HTTP on this host:port")
	flag.Parse()
	wd, _ := filepath.Abs(*dir)
	os.Chdir(wd)
	log.Printf("Starting HTTP Server at %q", *httpl)
	h := simplehttp.New(wd, *dev)
	log.Fatal(http.ListenAndServe(*httpl, h))
}
