package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

var quotes = []string{
	"The only limit to our realization of tomorrow is our doubts of today.",
	"The purpose of our lives is to be happy.",
	"Life is what happens when you're busy making other plans.",
	"Get busy living or get busy dying.",
}

func generateProofOfWork(challenge string) (string, int) {
	var nonce int
	var hash [32]byte
	for {
		nonce = rand.Int()
		data := fmt.Sprintf("%s%d", challenge, nonce)
		hash = sha256.Sum256([]byte(data))
		if strings.HasPrefix(hex.EncodeToString(hash[:]), "0000") {
			break
		}
	}
	return hex.EncodeToString(hash[:]), nonce
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	rand.Seed(time.Now().UnixNano())
	challenge := strconv.Itoa(rand.Intn(1000000))

	fmt.Println("Generated challenge:", challenge)
	conn.Write([]byte(challenge + "\n"))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	nonceStr := strings.TrimSpace(string(buf[:n]))
	nonce, err := strconv.Atoi(nonceStr)
	if err != nil {
		fmt.Println("Invalid nonce:", nonceStr)
		return
	}

	data := fmt.Sprintf("%s%d", challenge, nonce)
	hash := sha256.Sum256([]byte(data))
	if strings.HasPrefix(hex.EncodeToString(hash[:]), "0000") {
		quote := quotes[rand.Intn(len(quotes))]
		conn.Write([]byte(quote + "\n"))
	} else {
		conn.Write([]byte("Invalid PoW\n"))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Error setting up server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server listening on port 9999")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
