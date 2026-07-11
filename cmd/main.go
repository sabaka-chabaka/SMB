package main

import (
	"SMB/internal/broker"
	"SMB/internal/config"
	"log"
	"net/http"
)

func main() {
	config.Load()

	brk := broker.NewBroker()

	http.HandleFunc("/ws", broker.Handler(brk))

	log.Printf("Server started on %s\n", config.Configuration.SMBPort)
	if err := http.ListenAndServe(config.Configuration.SMBPort, nil); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
