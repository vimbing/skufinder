package internal

import http "github.com/vimbing/fhttp"

func (f *Finder) Headers() http.Header {
	return http.Header{
		"authority":                  {"lens.google.com"},
		"accept":                     {"*/*"},
		"accept-language":            {"en-US,en"},
		"content-type":               {"application/x-www-form-urlencoded;charset=utf-8"},
		"dnt":                        {"1"},
		"sec-ch-ua-arch":             {"\"x86\""},
		"sec-ch-ua-bitness":          {"\"64\""},
		"sec-ch-ua-full-version":     {"\"113.0.5672.126\""},
		"sec-ch-ua-mobile":           {"?0"},
		"sec-ch-ua-model":            {"\"\""},
		"sec-ch-ua-platform-version": {"\"5.15.0\""},
		"sec-ch-ua-wow64":            {"?0"},
		"sec-fetch-dest":             {"empty"},
		"sec-fetch-mode":             {"cors"},
		"sec-fetch-site":             {"same-origin"},
		"user-agent":                 {"Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/116.0.5845.118 Mobile/15E148 Safari/604.1"},
		"x-client-side-image-upload": {"true"},
		"x-goog-upload-command":      {"upload, finalize"},
		"x-goog-upload-offset":       {"0"},
	}
}

func (f *Finder) SearchHeaders() http.Header {
	return http.Header{
		"authority":                   {"lens.google.com"},
		"accept":                      {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"accept-language":             {"en-US,en"},
		"dnt":                         {"1"},
		"referer":                     {"https://www.google.pl/"},
		"sec-ch-ua":                   {"\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\""},
		"sec-ch-ua-arch":              {"\"x86\""},
		"sec-ch-ua-bitness":           {"\"64\""},
		"sec-ch-ua-full-version":      {"\"113.0.5672.126\""},
		"sec-ch-ua-full-version-list": {"\"Google Chrome\";v=\"113.0.5672.126\", \"Chromium\";v=\"113.0.5672.126\", \"Not-A.Brand\";v=\"24.0.0.0\""},
		"sec-ch-ua-mobile":            {"?0"},
		"sec-ch-ua-model":             {"\"\""},
		"sec-ch-ua-platform":          {"\"Linux\""},
		"sec-ch-ua-platform-version":  {"\"5.15.0\""},
		"sec-ch-ua-wow64":             {"?0"},
		"sec-fetch-dest":              {"document"},
		"sec-fetch-mode":              {"navigate"},
		"sec-fetch-site":              {"same-origin"},
		"upgrade-insecure-requests":   {"1"},
		"user-agent":                  {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"},
	}

}

func (f *Finder) AddAdditionalCheckFunc(fc AdditionalCheckFunc) {
	f.AdditionalCheckFunc = fc
}
