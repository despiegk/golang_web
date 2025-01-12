package main

import (
	"heroweb/internal/dns"
	"heroweb/internal/filemanager"
	"heroweb/internal/server"
	"log"
)

func main() {

	server := server.NewServer()

	// Create a channel to listen for errors
	//errChan := make(chan error)

	// Start the server in a goroutine
	// go func() {
	// 	errChan <- server.ListenAndServe()
	// }()

	// Initialize file manager with current directory as root
	_ = filemanager.NewFileManager(".")

	// Start DNS services
	dns.DNSDBStart()
	if err := dns.RegisterARecord("mysite.mmm", "1.1.1.1"); err != nil {
		log.Printf("Warning: Failed to register DNS records: %v", err)
	}
	go dns.DnsStart()

	// Start the server (this will block)
	log.Fatal(server.ListenAndServe())
}
