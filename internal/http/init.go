package http

import tls "github.com/vimbing/utls"

func Init(opts ...Options) (*Client, error) {
	client := Client{}

	if len(opts) > 0 {
		client.Hello = opts[0].Hello
	} else {
		client.Hello = tls.HelloChrome_100
	}

	err := client.InitClient()

	return &client, err
}
