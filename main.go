package main

import (
	"log"
	"net/url"
	"os"
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

func parseURL(input string) *url.URL {
	url, err := url.Parse(input)
	if err != nil {
		log.Fatal("Problem parsing input: ", input)
	}
	if url.Scheme == "" {
		log.Println("Scheme was not given, assuming scheme gemini")
		url.Scheme = "gemini"
	}
	// Spec Note:
	// 	Sending an absolute URL instead of only a path or selector is
	// effectively equivalent to building in a HTTP "Host" header.  It
	// permits virtual hosting of multiple Gemini domains on the same IP
	// address.  It also allows servers to optionally act as proxies.
	// Including schemes other than gemini:// in requests allows servers to
	// optionally act as protocol-translating gateways to e.g. fetch gopher
	// resources over Gemini.  Proxying is optional and the vast majority of
	// servers are expected to only respond to requests for resources at
	// their own domain(s).
	return url
}

func main() {
	args := os.Args
	url := parseURL(args[1])
	log.Println("Visiting ", url)

}
