package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ngynkvn/gemgo/gemini"
)

func GeminiTransaction() {
	//C:   Opens connection
	//S:   Accepts connection
	//C/S: Complete TLS handshake (see 1.4)
	//C:   Validates server certificate (see 1.4.2)
	//C:   Sends request (one CRLF terminated line) (see 1.2)
	//S:   Sends response header (one CRLF terminated line), closes connection
	//     under non-success conditions (see 1.3.1, 1.3.2)
	//S:   Sends response body (text or binary data) (see 1.3.3)
	//S:   Closes connection
	//C:   Handles response (see 1.3.4)
}

func main() {
	args := os.Args
	url := gemini.ParseURL(args[1])
	log.Println("Visiting", url.String())

	gc := gemini.NewGeminiConnection(url)

	gc.SendRequest(url)
	header := gc.ReceiveHeader()
	fmt.Println("header", header)
	gc.ReceiveBody()
	// fmt.Println(body)
}
