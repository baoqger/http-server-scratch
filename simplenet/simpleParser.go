package simplenet

import (
	"bufio"
	"errors"
	"io"
	"net/textproto"
	"strconv"
	"strings"
)

type request struct {
	method string // GET, POST, etc.
	header textproto.MIMEHeader
	body   []byte
	uri    string // The raw URI from the request
	proto  string // "HTTP/1.1"
}

func ParseRequest(c *netSocket) (*request, error) {
	b := bufio.NewReader(*c)
	tp := textproto.NewReader(b)
	req := new(request)

	// First line: parse "GET /index.html HTTP/1.0"
	var s string
	s, _ = tp.ReadLine()
	sp := strings.Split(s, " ")
	req.method, req.uri, req.proto = sp[0], sp[1], sp[2]

	// Parse headers
	mimeHeader, _ := tp.ReadMIMEHeader()
	req.header = mimeHeader

	// Parse body
	if req.method == "GET" || req.method == "HEAD" {
		return req, nil
	}
	if len(req.header["Content-Length"]) == 0 {
		return nil, errors.New("no content length")
	}
	length, err := strconv.Atoi(req.header["Content-Length"][0])
	if err != nil {
		return nil, err
	}
	body := make([]byte, length)
	if _, err = io.ReadFull(b, body); err != nil {
		return nil, err
	}
	req.body = body
	return req, nil
}
