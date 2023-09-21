package api

type JavaScript struct {
	Func        string `json:"func"`
	ShouldParse bool   `json:"shouldParse"`
}

type Config struct {
	MinimumLength int        `json:"minimumLength"`
	MaximumLength int        `json:"maximumLength"`
	JavaScript    JavaScript `json:"javaScript"`
	Regexp        string     `json:"regexp"`
}

type Payload struct {
	Image  string `json:"image"`
	Config Config `json:"config"`
}
