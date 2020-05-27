package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ngynkvn/gemgo/gemini"
)

func main() {
	args := os.Args
	url := gemini.ParseURL(args[1])
	log.Println("Visiting", url.String())

	gc := gemini.NewGeminiConnection(url)

	gc.SendRequest(url)
	header := gc.ReceiveHeader()
	fmt.Println("header", header)
	body := gc.ReceiveBody()
	fmt.Println(body)
}
