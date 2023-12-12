package finder

import (
	"github.com/vimbing/http"

	tls "github.com/vimbing/utls"
)

func Init(photo string, cfg *Config) *Finder {
	return &Finder{
		internal: internal{
			Client:       http.Init(http.Config{TlsHello: &tls.HelloIOS_12_1}),
			WordsToCheck: make([]SkuFinderWord, 0),
			Occurences:   make(map[string]int),
			config: config{
				Photo:               photo,
				SkuRegexp:           cfg.SkuRegexp,
				AdditionalCheckFunc: cfg.AdditionalCheckFunc,
				MinimumLength:       cfg.MinimumLength,
				MaximumLength:       cfg.MaximumLength,
			},
		},
	}
}
