package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"strings"
)

func solveProofOfWork(challenge string) int {
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
	return nonce
}

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	message, _ := bufio.NewReader(conn).ReadString('\n')
	challenge := strings.TrimSpace(message)

	nonce := solveProofOfWork(challenge)
	fmt.Fprintf(conn, "%d\n", nonce)

	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Received response: ", response)
}
