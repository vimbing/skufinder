package internal

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"skufinder/internal/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/percolate/retry"
)

func (f *Finder) getPhoto() error {
	fmt.Println("Getting photo bytes...")

	return retry.Re{Max: 5, Delay: time.Second * 3}.Try(func() error {
		res, err := f.internal.Client.Request(&http.Request{
			Method: "GET",
			Url:    f.Photo,
		})

		if err != nil {
			return err
		}

		defer res.Body.Close()

		photoBytes, err := f.internal.Client.GetResponseBodyBytes(res)

		if err != nil {
			return err
		}

		f.internal.PhotoBytes = photoBytes

		fmt.Println("Photo bytes successfully scraped...")

		return nil
	})
}

func (f *Finder) uploadPhoto() error {
	fmt.Println("Uploading photo to google lens...")

	return retry.Re{Max: 5, Delay: time.Second * 3}.Try(func() error {
		res, err := f.internal.Client.Request(&http.Request{
			Method:  "POST",
			Url:     "https://lens.google.com/_/upload/",
			Body:    bytes.NewBuffer(f.internal.PhotoBytes),
			Headers: f.Headers(),
		})

		if err != nil {
			return err
		}

		defer res.Body.Close()

		body, err := f.internal.Client.GetResponseBodyString(res)

		if err != nil {
			return err
		}

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
		res, err := f.internal.Client.Request(&http.Request{
			Method:  "GET",
			Url:     f.internal.Url,
			Headers: f.SearchHeaders(),
		})

		if err != nil {
			return err
		}

		defer res.Body.Close()

		f.internal.LensResultBody, err = f.internal.Client.GetResponseBodyString(res)

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
			if len(word) >= f.MinimumLength && len(word) <= f.MaximumLength && f.SkuRegexp.Match([]byte(word)) && !regexp.MustCompile("[!@#$%^&*()_+=\\[\\]{};':\"\\\\|,.<>/?`~]").MatchString(word) {
				word = f.SkuRegexp.FindString(word)

				if f.AdditionalCheckFunc != nil && !f.AdditionalCheckFunc(word) {
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
		return errors.New("no_words")
	}

	for word, count := range words {
		f.internal.WordsToCheck = append(f.internal.WordsToCheck, SkuFinderWord{
			Count:     count,
			Word:      word,
			StockxUrl: fmt.Sprintf("https://stockx.com/search?s=%s", word),
		})
	}

	return nil
}

func (f *Finder) findSkuFromOtherString() error {
	words := make(map[string]int)

	for _, word := range strings.Split(f.internal.OtherString, " ") {
		if len(word) >= f.MinimumLength && len(word) <= f.MaximumLength && f.SkuRegexp.Match([]byte(word)) && !regexp.MustCompile("[!@#$%^&*()_+=\\[\\]{};':\"\\\\|,.<>/?`~]").MatchString(word) {
			word = f.SkuRegexp.FindString(word)

			if f.AdditionalCheckFunc != nil && !f.AdditionalCheckFunc(word) {
				continue
			}

			if _, ok := words[strings.ToLower(word)]; ok {
				words[strings.ToLower(word)]++
				continue
			}

			words[strings.ToLower(word)] = 1
		}
	}

	if len(words) == 0 {
		return errors.New("no_words")
	}

	for word, count := range words {
		f.internal.WordsToCheck = append(f.internal.WordsToCheck, SkuFinderWord{
			Count:     count,
			Word:      word,
			StockxUrl: fmt.Sprintf("https://stockx.com/search?s=%s", word),
		})
	}

	return nil
}

func (f *Finder) Sort() ([]SkuFinderWord, error) {
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
		return []SkuFinderWord{}, err
	}

	return f.Sort()
}

func (f *Finder) GetSkuFromString(otherString string) ([]SkuFinderWord, error) {
	f.internal.OtherString = otherString

	err := f.findSkuFromOtherString()

	if err != nil {
		return []SkuFinderWord{}, err
	}

	return f.Sort()
}
