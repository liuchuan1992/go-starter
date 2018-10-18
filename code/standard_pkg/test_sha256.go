package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	data := []byte("12334531")
	hash := sha256.New()
	hash.Write(data)
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	fmt.Println(mdStr)
}