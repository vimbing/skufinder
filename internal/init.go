package internal

import (
	"skufinder/internal/http"

	tls "github.com/vimbing/utls"
)

func Init(photo string, config *Config) (*Finder, error) {
	client, err := http.Init(http.Options{Hello: tls.HelloIOS_12_1})

	if err != nil {
		return &Finder{}, err
	}

	return &Finder{
		internal: internal{
			Client:       client,
			WordsToCheck: make([]SkuFinderWord, 0),
			Occurences:   make(map[string]int),
		},
		Photo:               photo,
		SkuRegexp:           config.SkuRegexp,
		AdditionalCheckFunc: config.AdditionalCheckFunc,
		MinimumLength:       config.MinimumLength,
		MaximumLength:       config.MaximumLength,
	}, nil
}
