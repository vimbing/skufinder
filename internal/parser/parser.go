package parser

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func (p *Parser) findAfInitDataPart() error {
	afInitDataRe := regexp.MustCompile(`>AF_initDataCallback[^<]+`)
	finds := afInitDataRe.FindAllString(string(p.body), -1)

	var usedScript string

	for _, find := range finds {
		if len(usedScript) < len(find) {
			usedScript = find
		}
	}

	if len(usedScript) < 3 || len(usedScript) < 21 {
		return ErrWrongAfInitFind
	}

	p.afInitDataPart = usedScript[21 : len(usedScript)-2]

	return nil
}

func (p *Parser) parseBody() {
	cleanRe1 := regexp.MustCompile(`,\ssideChannel:\s{}`)
	cleanRe2 := regexp.MustCompile(`key:\s'ds:\d+',\shash:\s'\d+',\s`)

	p.afInitDataPart = cleanRe1.ReplaceAllString(p.afInitDataPart, "")
	p.afInitDataPart = cleanRe2.ReplaceAllString(p.afInitDataPart, "")

	p.afInitDataPart = strings.Replace(p.afInitDataPart, "data:", "\"data\":", -1)
}

func (p *Parser) unmarshalCleanedAfInitData() {
	json.Unmarshal([]byte(p.afInitDataPart), &p.parsedAfInitData)
}

func parseItem(data []any) Item {
	thumbnailImage := data[0].([]interface{})[0]
	title := data[1].(string)
	description := data[2].(string)
	url := data[5].(string)

	return Item{
		ThumbnailImg: thumbnailImage.(string),
		Title:        title,
		Description:  description,
		Url:          url,
	}
}

func (p *Parser) getParsedItems() []Item {
	parsedItems := []Item{}

	for _, item := range p.parsedAfInitData.Data[1][0][1][8][8][0][12][0][9][0] {
		parsedItems = append(parsedItems, parseItem(item))
	}

	return parsedItems
}

func (p *Parser) Parse() ([]Item, error) {
	if err := p.findAfInitDataPart(); err != nil {
		fmt.Println(err)
		return []Item{}, err
	}

	p.parseBody()
	p.unmarshalCleanedAfInitData()

	return p.getParsedItems(), nil
}
