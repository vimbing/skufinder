package http

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/andybalholm/brotli"
	"github.com/vimbing/cclient"
	http "github.com/vimbing/fhttp"
)

func (c *Client) Decode(res *http.Response, out any) error {
	bodyBytes, err := c.GetResponseBodyBytes(res)

	if err != nil {
		return err
	}

	json.Unmarshal(bodyBytes, out)

	return nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.Client.Do(req)
}

func (c *Client) Request(r *Request) (*http.Response, error) {
	req, err := http.NewRequest(r.Method, r.Url, r.Body)

	if err != nil {
		pw, _ := io.Pipe()
		return &http.Response{Body: pw}, err
	}

	req.Header = r.Headers

	return c.Do(req)
}

func (c *Client) InitClient() error {
	var jarCopy http.CookieJar
	if c.Client != nil {
		jarCopy = c.Client.Jar
	}

	client, err := cclient.NewClient(c.Hello, c.Proxy, true, 15)

	if err != nil {
		return err
	}

	client.Jar = jarCopy

	c.Client = &client

	return nil
}

func (c *Client) ChangeProxy(p string) error {
	c.Proxy = p
	return c.InitClient()
}

func (c *Client) GetResponseBodyString(r *http.Response) (string, error) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return "", err
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	decodedBody, err := DecodeFhttp(r.Header, body)

	if err != nil {
		return "", err
	}

	return decodedBody, nil
}

func (c *Client) GetResponseBodyBytes(r *http.Response) ([]byte, error) {
	bodyString, err := c.GetResponseBodyString(r)
	return []byte(bodyString), err
}

func DecodeFhttp(headers http.Header, body []byte) (string, error) {
	defer func() (string, error) {
		if err := recover(); err != nil {
			return "", errors.New("asd")
		}
		return "", nil
	}()

	var encoding string

	if len(headers["Content-Encoding"]) == 0 {
		encoding = "NAN"
	} else {
		encoding = headers["Content-Encoding"][0]
	}

	if encoding == "br" {
		decodedBody, err := unBrotliData(body)

		if err != nil {
			return "", err
		}

		return string(decodedBody), nil
	} else if encoding == "deflate" {
		decodedBody, err := enflateData(body)

		if err != nil {
			return "", err
		}

		return string(decodedBody), nil
	} else if encoding == "gzip" {
		decodedBody, err := gUnzipData(body)

		if err != nil {
			return "", err
		}

		return string(decodedBody), nil
	} else {
		return (string(body)), nil
	}
}

func unBrotliData(data []byte) (resData []byte, err error) {
	br := brotli.NewReader(bytes.NewReader(data))
	respBody, err := io.ReadAll(br)
	return respBody, err
}

func enflateData(data []byte) (resData []byte, err error) {
	zr, _ := zlib.NewReader(bytes.NewReader(data))
	defer zr.Close()
	enflated, err := io.ReadAll(zr)
	return enflated, err
}

func gUnzipData(data []byte) (resData []byte, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("")
		}
	}()
	gz, _ := gzip.NewReader(bytes.NewReader(data))
	defer gz.Close()
	respBody, err := io.ReadAll(gz)
	return respBody, err
}
