package internal

import (
	"regexp"
	"skufinder/internal/http"
)

type internal struct {
	Client         *http.Client
	Url            string
	LensResultBody string
	WordsToCheck   []SkuFinderWord
	Occurences     map[string]int
	PhotoBytes     []byte
	OtherString    string
}

type AdditionalCheckFunc func(string) bool

type Finder struct {
	internal            internal
	SkuRegexp           *regexp.Regexp
	AdditionalCheckFunc AdditionalCheckFunc
	MinimumLength       int
	MaximumLength       int
	Photo               string
}

type Config struct {
	SkuRegexp           *regexp.Regexp
	MinimumLength       int
	MaximumLength       int
	AdditionalCheckFunc AdditionalCheckFunc
}

type SkuFinderWord struct {
	Count     int    `json:"count"`
	Word      string `json:"word"`
	StockxUrl string `json:"stockxUrl"`
}
