package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("Please provide url to encode")
	}

	url := args[1]
	filename := "qrcode.png"
	err := qrcode.WriteFile(url, qrcode.Medium, 256, filename)
	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Println("QR code generated and saved as:", filename)
}
