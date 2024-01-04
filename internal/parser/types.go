package parser

type AFinitData struct {
	Data [][][][][][][][][][][][]any
}

type Item struct {
	ThumbnailImg string
	Title        string
	Description  string
	Url          string
}

type Parser struct {
	body             string
	afInitDataPart   string
	parsedAfInitData AFinitData
}
