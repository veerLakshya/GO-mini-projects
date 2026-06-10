package main

import (
	"log"
	"time"
)

func main() {
	server := NewFileServer(":4040")

	// START SERVER
	go func() {
		if err := server.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	// GIVE SERVER TIME TO START
	time.Sleep(2 * time.Second)

	// SEND FILE
	err := SendFile(":4040", 10*1024*1024) // 10 MB
	if err != nil {
		log.Fatal(err)
	}

	// KEEP PROCESS ALIVE
	select {}
}
