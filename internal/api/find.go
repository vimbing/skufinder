package api

import (
	"encoding/base64"
	"regexp"
	"skufinder/internal"

	"github.com/gofiber/fiber/v2"
)

func Find(c *fiber.Ctx) error {
	var payload Payload
	c.BodyParser(&payload)

	re, err := base64.StdEncoding.DecodeString(payload.Config.Regexp)

	if err != nil {
		return c.SendStatus(503)
	}

	var additionalCheckFunc internal.AdditionalCheckFunc
	additionalCheckFunc = nil

	if payload.Config.JavaScript.ShouldParse {
		jsFunc, err := base64.StdEncoding.DecodeString(payload.Config.JavaScript.Func)

		if err != nil {
			return c.SendStatus(503)
		}

		additionalCheckFunc = getAdditionalCheckFuncFromJs(string(jsFunc))
	}

	config := internal.Config{
		SkuRegexp:           regexp.MustCompile(string(re)),
		MinimumLength:       payload.Config.MinimumLength,
		MaximumLength:       payload.Config.MaximumLength,
		AdditionalCheckFunc: additionalCheckFunc,
	}

	fnder, err := internal.Init(payload.Image, &config)

	if err != nil {
		return c.SendStatus(503)
	}

	words, err := fnder.GetSku()

	if err != nil {
		return c.SendStatus(503)
	}

	return c.JSON(words)
}
