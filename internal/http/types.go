package http

import (
	"io"

	http "github.com/vimbing/fhttp"
	tls "github.com/vimbing/utls"
)

type Client struct {
	Hello  tls.ClientHelloID
	Client *http.Client
	Proxy  string
}

type Request struct {
	Method    string
	Url       string
	Body      io.Reader
	Headers   http.Header
	Proxyless bool
}

type Options struct {
	Hello tls.ClientHelloID
}
