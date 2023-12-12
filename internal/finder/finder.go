package finder

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/percolate/retry"
)

func (f *Finder) getPhoto() error {
	fmt.Println("Getting photo bytes...")

	return retry.Re{Max: 5, Delay: time.Second * 3}.Try(func() error {

		res, err := f.internal.Client.Get(f.internal.config.Photo)

		if err != nil {
			return err
		}

		f.internal.PhotoBytes = res.Body

		fmt.Println("Photo bytes successfully scraped...")

		return nil
	})
}

func (f *Finder) uploadPhoto() error {
	fmt.Println("Uploading photo to google lens...")

	return retry.Re{Max: 5, Delay: time.Second * 3}.Try(func() error {
		res, err := f.internal.Client.Post(
			"https://lens.google.com/_/upload/",
			bytes.NewBuffer(f.internal.PhotoBytes),
			f.headers(),
		)

		if err != nil {
			return err
		}

		body := res.BodyString()

		urlFirst := regexp.MustCompile(`":"\/[^"]+`).FindString(string(body))

		if len(urlFirst) <= 3 {
			return errors.New("wrong_url")
		}

		decodedUrl, err := strconv.Unquote(`"` + urlFirst[3:] + `"`)

		if err != nil {
			return err
		}

		f.internal.Url = fmt.Sprintf("https://lens.google.com%s", decodedUrl)

		fmt.Println("Photo successfully uploaded to google lens!")

		return nil
	})
}

func (f *Finder) getLensResults() error {
	fmt.Println("Scraping google lens results...")

	return retry.Re{Max: 5, Delay: time.Second * 3}.Try(func() error {
		res, err := f.internal.Client.Get(
			f.internal.Url,
			f.searchHeaders(),
		)

		if err != nil {
			return err
		}

		f.internal.LensResultBody = res.BodyString()

		fmt.Println("Google lens results successfully scraped!")

		return err
	})
}

func (f *Finder) findSku() error {
	resultDatasRegexp := regexp.MustCompile(`>AF_initDataCallback[^<]+`)
	resultDatas := resultDatasRegexp.FindAllString(f.internal.LensResultBody, -1)

	words := make(map[string]int)

	for _, resultData := range resultDatas {
		for _, word := range strings.Split(resultData, " ") {
			if len(word) >= f.internal.config.MinimumLength && len(word) <= f.internal.config.MaximumLength && f.internal.config.SkuRegexp.Match([]byte(word)) && !regexp.MustCompile("[!@#$%^&*()_+=\\[\\]{};':\"\\\\|,.<>/?`~]").MatchString(word) {
				word = f.internal.config.SkuRegexp.FindString(word)

				if f.internal.config.AdditionalCheckFunc != nil && !f.internal.config.AdditionalCheckFunc(word) {
					continue
				}

				if _, ok := words[strings.ToLower(word)]; ok {
					words[strings.ToLower(word)]++
					continue
				}

				words[strings.ToLower(word)] = 1
			}
		}
	}

	if len(words) == 0 {
		return ErrNoWords
	}

	for word, count := range words {
		f.internal.WordsToCheck = append(f.internal.WordsToCheck, SkuFinderWord{
			Count: count,
			Word:  word,
		})
	}

	return nil
}

func (f *Finder) sort() ([]SkuFinderWord, error) {
	sort.Slice(f.internal.WordsToCheck, func(i, j int) bool {
		return f.internal.WordsToCheck[i].Count > f.internal.WordsToCheck[j].Count
	})

	if len(f.internal.WordsToCheck) <= 5 {
		return f.internal.WordsToCheck, nil
	}

	return f.internal.WordsToCheck[:5], nil
}

func (f *Finder) GetSku() ([]SkuFinderWord, error) {
	err := f.getPhoto()

	if err != nil {
		return []SkuFinderWord{}, err
	}

	err = f.uploadPhoto()

	if err != nil {
		return []SkuFinderWord{}, err
	}

	err = f.getLensResults()

	if err != nil {
		return []SkuFinderWord{}, err
	}

	err = f.findSku()

	if err != nil {
		if err == ErrNoWords {
			return []SkuFinderWord{}, nil
		}

		return []SkuFinderWord{}, err
	}

	return f.sort()
}
