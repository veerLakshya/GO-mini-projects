package main

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"os"
)

type FileServer struct {
	addr string
}

func NewFileServer(addr string) *FileServer {
	return &FileServer{
		addr: addr,
	}
}

func (fs *FileServer) Start() error {
	ln, err := net.Listen("tcp", fs.addr)
	if err != nil {
		return err
	}

	log.Println("server listening on", fs.addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}

		go fs.handleConn(conn)
	}
}

func (fs *FileServer) handleConn(conn net.Conn) {
	defer conn.Close()

	log.Println("new connection:", conn.RemoteAddr())

	// -----------------------------
	// READ FILE SIZE HEADER
	// -----------------------------

	var fileSize uint64

	err := binary.Read(conn, binary.LittleEndian, &fileSize)
	if err != nil {
		log.Println("failed to read file size:", err)
		return
	}

	log.Println("incoming file size:", fileSize)

	// -----------------------------
	// CREATE DESTINATION FILE
	// -----------------------------

	file, err := os.Create("received.bin")
	if err != nil {
		log.Println("failed to create file:", err)
		return
	}
	defer file.Close()

	// -----------------------------
	// COPY FILE CONTENT
	// -----------------------------

	n, err := io.CopyN(file, conn, int64(fileSize))
	if err != nil {
		log.Println("failed to receive file:", err)
		return
	}

	log.Println("received", n, "bytes successfully")
}
