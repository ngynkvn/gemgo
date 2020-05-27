package gemini

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Status int
type LineType int

// Status Codes
const (
	Invalid                       Status = -1
	Input                         Status = 10
	Success                       Status = 20
	RedirectTemp                  Status = 30
	RedirectPermanent             Status = 31
	TemporaryFailure              Status = 40
	ServerUnavailable             Status = 41
	CgiError                      Status = 42
	ProxyError                    Status = 43
	Slowdown                      Status = 44
	PermanentFailure              Status = 50
	NotFound                      Status = 51
	Gone                          Status = 52
	ProxyRequestRefused           Status = 53
	BadRequest                    Status = 59
	ClientCertificateRequired     Status = 60
	TransientCertificateRequested Status = 61
	AuthorisedCertificateRequired Status = 62
	CertificateNotAccepted        Status = 63
	FutureCertificateRejected     Status = 64
	ExpiredCertificateRejected    Status = 65
)

// LineTypes ...
const (
	Text LineType = iota
	Link
	PreformatToggle
	// TODO Advanced Lines
)

var statusCodes = []Status{
	Input,
	Success,
	RedirectTemp,
	RedirectPermanent,
	TemporaryFailure,
	ServerUnavailable,
	CgiError,
	ProxyError,
	Slowdown,
	PermanentFailure,
	NotFound,
	Gone,
	ProxyRequestRefused,
	BadRequest,
	ClientCertificateRequired,
	TransientCertificateRequested,
	AuthorisedCertificateRequired,
	CertificateNotAccepted,
	FutureCertificateRejected,
	ExpiredCertificateRejected,
}

type GeminiConnection struct {
	tlsConnection *tls.Conn
	scanner       *bufio.Scanner
}

type Header struct {
	status Status
	meta   string
	raw    string
}

func (header Header) String() string {
	return fmt.Sprintf("[%d] %s", header.status, header.meta)
}

type Line struct {
	lineType LineType
	meta     string
	raw      string
}

type Body struct {
	Lines []Line
}

func (body Body) String() string {
	var str bytes.Buffer
	for i, v := range body.Lines {
		// Print raw string for now.
		str.WriteString(fmt.Sprintf("%d\t| %s\n", i, v.raw))
	}
	return str.String()
}

const TERMINATOR = "\r\n"

func NewGeminiConnection(url URL) GeminiConnection {
	conn, err := url.Dial()
	if err != nil {
		log.Fatal("Problem connecting to url: ", url.Addr())
	}
	scanner := bufio.NewScanner(conn)
	return GeminiConnection{conn, scanner}
}

func (gc *GeminiConnection) SendRequest(url URL) (int, error) {
	n, err := gc.tlsConnection.Write([]byte(url.String() + TERMINATOR))
	return n, err
}

func interpretHeader(header string) Header {
	status, meta := matchStatus(header)
	return Header{status, meta, header}
}

func matchStatus(header string) (Status, string) {
	for _, status := range statusCodes {
		if strings.HasPrefix(header, strconv.Itoa(int(status))) {
			return Status(status), header[3:]
		}
	}
	return Invalid, ""
}

func (gc *GeminiConnection) ReceiveHeader() Header {
	if !gc.scanner.Scan() {
		if err := gc.scanner.Err(); err != nil {
			panic(err)
		}
	}
	return interpretHeader(gc.scanner.Text())
}

func interpretBody(lines []Line) string {
	for i, v := range lines {
		println(i, v.raw)
	}
	return ""
}

func parseLine(text string) Line {
	switch {
	case strings.HasPrefix(text, "=> "):
		link := text[3:]
		return Line{Link, link, text}
	case text == "```":
		return Line{PreformatToggle, "", text}
	default:
		return Line{Text, "", text}
	}
}

func (gc *GeminiConnection) readBodyLines() []Line {
	// Scan entire body
	var lines []Line
	for gc.scanner.Scan() {
		text := gc.scanner.Text()
		lines = append(lines, parseLine(text))
	}
	if err := gc.scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func (gc *GeminiConnection) ReceiveBody() Body {
	lines := gc.readBodyLines()
	return Body{lines}
}
