package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"log"
	"net"
)

func SendFile(addr string, size int) error {
	// -----------------------------
	// GENERATE RANDOM FILE DATA
	// -----------------------------

	fileData := make([]byte, size)

	_, err := io.ReadFull(rand.Reader, fileData)
	if err != nil {
		return err
	}

	// -----------------------------
	// CONNECT TO SERVER
	// -----------------------------

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	log.Println("connected to server")

	// -----------------------------
	// SEND FILE SIZE HEADER
	// -----------------------------

	err = binary.Write(conn, binary.LittleEndian, uint64(len(fileData)))
	if err != nil {
		return err
	}

	// -----------------------------
	// SEND FILE CONTENT
	// -----------------------------

	n, err := io.Copy(conn, bytes.NewReader(fileData))
	if err != nil {
		return err
	}

	log.Println("sent", n, "bytes successfully")

	return nil
}
