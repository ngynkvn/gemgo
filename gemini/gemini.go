package gemini

import (
	"crypto/tls"
	"io"
	"log"
)

type Status int

// Status Codes
const (
	INPUT   Status = 10
	SUCCESS Status = 20
)

type GeminiConnection struct {
	tlsConnection *tls.Conn
}

type Header struct {
	raw string
}

const TERMINATOR = "\r\n"

func NewGeminiConnection(url URL) GeminiConnection {
	conn, err := url.Dial()
	if err != nil {
		log.Fatal("Problem connecting to url: ", url.Addr())
	}
	return GeminiConnection{conn}
}

func (gc *GeminiConnection) SendRequest(url URL) (int, error) {
	n, err := gc.tlsConnection.Write([]byte(url.String() + TERMINATOR))
	return n, err
}

func interpretHeader(header string) Header {
	//TODO, extract actual info
	return Header{header}
}

func (gc *GeminiConnection) ReceiveHeader() Header {
	buffer := make([]byte, 256) // TODO Possible buffer overflow?
	//TODO Deadlines
	if n, err := gc.tlsConnection.Read(buffer); err != nil {
		log.Fatal(n, err)
	}
	header := string(buffer)
	return interpretHeader(header)
}

func (gc *GeminiConnection) ReceiveBody() string {
	buffer := make([]byte, 1024) // TODO Possible buffer overflow?
	var body string
	//TODO Deadlines
	for {
		n, err := gc.tlsConnection.Read(buffer)
		body = body + string(buffer[:n])
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Err", n, err)
		}
	}
	return body
}
