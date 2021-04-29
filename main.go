package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		flagAddr   string
		flagFolder string
	)
	log.SetFlags(0)
	flag.StringVar(&flagAddr, "addr", ":8080", "IP:Port to use for the webserver")
	flag.StringVar(&flagFolder, "folder", "/var/www/", "Folder to serve static content from")
	// parse cmd-line parameters
	flag.Usage = func() {
		fmt.Println("This is a dead-simple webserver serving (static) contents of the specified folder on the specified (http-only) address.")
		fmt.Println("This program is maintained at https://github.com/prigio/gotpd")
		fmt.Println("")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if _, err := os.Stat(flagFolder); os.IsNotExist(err) {
		log.Fatalf("Configured folder does not exist '%s'\n", flagFolder)
		os.Exit(1)
	}

	log.Printf("Serving folder '%s' on address '%s'\n", flagFolder, flagAddr)
	log.Fatal(http.ListenAndServe(flagAddr, http.FileServer(http.Dir(flagFolder))))
}
