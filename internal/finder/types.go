package finder

import (
	"regexp"

	"github.com/vimbing/http"
)

type config struct {
	SkuRegexp           *regexp.Regexp
	AdditionalCheckFunc AdditionalCheckFunc
	MinimumLength       int
	MaximumLength       int
	Photo               string
}

type internal struct {
	Client         *http.Client
	Url            string
	LensResultBody string
	WordsToCheck   []SkuFinderWord
	Occurences     map[string]int
	PhotoBytes     []byte
	OtherString    string
	config         config
}

type AdditionalCheckFunc func(string) bool

type Finder struct {
	internal internal
}

type Config struct {
	SkuRegexp           *regexp.Regexp
	MinimumLength       int
	MaximumLength       int
	AdditionalCheckFunc AdditionalCheckFunc
}

type SkuFinderWord struct {
	Count int    `json:"count"`
	Word  string `json:"word"`
}
