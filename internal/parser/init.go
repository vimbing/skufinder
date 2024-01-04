package parser

func Init(body string) *Parser {
	return &Parser{
		body: body,
	}
}
