package finder

import (
	"regexp"
	"unicode"
)

var (
	ConfigNike = Config{
		SkuRegexp:           regexp.MustCompile(`\b[\d\w]{3}[\d\w]{3}-[\d\w]{3}\b`),
		MinimumLength:       6,
		MaximumLength:       15,
		AdditionalCheckFunc: nil,
	}

	ConfigAsics = Config{
		SkuRegexp:     regexp.MustCompile(`^\d+[a-zA-Z]\d+-\d+$`),
		MinimumLength: 6,
		MaximumLength: 15,
		AdditionalCheckFunc: func(word string) bool {
			for _, char := range word {
				if unicode.IsLetter(char) {
					return true
				}
			}

			return false
		},
	}

	ConfigAdidas = Config{
		SkuRegexp:     regexp.MustCompile(`.*\d.*[a-zA-Z].*|.*[a-zA-Z].*\d.*`),
		MinimumLength: 5,
		MaximumLength: 15,
		AdditionalCheckFunc: func(s string) bool {
			wordRegexp := regexp.MustCompile(`[a-zA-Z]`)
			digitRegexp := regexp.MustCompile(`\d`)
			return wordRegexp.MatchString(s) && digitRegexp.MatchString(s)
		},
	}

	ConfigNewBalance = Config{
		SkuRegexp:     regexp.MustCompile(`.*\d.*[a-zA-Z].*|.*[a-zA-Z].*\d.*`),
		MinimumLength: 5,
		MaximumLength: 15,
		AdditionalCheckFunc: func(s string) bool {
			wordRegexp := regexp.MustCompile(`[a-zA-Z]`)
			digitRegexp := regexp.MustCompile(`\d`)
			return wordRegexp.MatchString(s) && digitRegexp.MatchString(s)
		},
	}

	ConfigBirkenstock = Config{
		SkuRegexp:           regexp.MustCompile(`^\d+$`),
		MinimumLength:       5,
		MaximumLength:       15,
		AdditionalCheckFunc: nil,
	}
)
