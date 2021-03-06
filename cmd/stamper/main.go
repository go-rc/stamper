package main

import (
	"flag"
	"log"
	"os"

	"github.com/tombell/stamper/services"
	"github.com/tombell/stamper/web"
)

func main() {
	host := flag.String("host", "127.0.0.1", "Host to bind to")
	port := flag.String("port", "8080", "Port to listen on")

	id := flag.String("id", "", "GitHub integration ID")
	cert := flag.String("cert", "", "Path to GitHub integration private key")

	needle := flag.String("needle", "", "String to find")
	label := flag.String("label", "bug", "Label to add to the issue/pull request")

	flag.Parse()

	l := log.New(os.Stderr, "[stamper] ", log.Lshortfile)
	l.Println("Starting...")

	var err error

	cfg := &services.GitHubServiceConfig{
		IntegrationID: *id,
		Cert:          *cert,
		Needle:        *needle,
		Label:         *label,
		Logger:        l,
	}

	err = services.SetupGitHubService(cfg)
	if err != nil {
		l.Fatal(err)
	}

	err = web.Run(*host, *port, l)
	if err != nil {
		l.Fatal(err)
	}
}
