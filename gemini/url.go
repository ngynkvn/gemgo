package gemini

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/url"
)

type URL struct {
	scheme string
	host   string
	path   string
	port   string
}

func ParseURL(input string) URL {
	url, err := url.Parse(input)
	if err != nil {
		log.Fatal("Problem parsing input: ", input)
	}
	var scheme string
	if scheme = url.Scheme; scheme == "" {
		log.Println("Scheme was not given, assuming scheme gemini")
		scheme = "gemini"
		url.Scheme = "gemini"
		url, err = url.Parse(url.String())
		if err != nil {
			log.Fatal("Problem parsing input: ", url.String())
		}
	}
	var port string
	if port = url.Port(); port == "" {
		log.Println("Log was not given, assuming port 1965")
		port = "1965"
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
	return URL{scheme, url.Host, url.Path, port}
}

func (url *URL) String() string {
	return fmt.Sprintf("%s://%s:%s%s", url.scheme, url.host, url.port, url.path)
}

func (url *URL) Addr() string {
	return fmt.Sprintf("%s:%s", url.host, url.port)
}

func (url *URL) Dial() (*tls.Conn, error) {
	return tls.Dial("tcp", url.Addr(), nil)
}
