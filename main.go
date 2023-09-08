package main

import (
	"fmt"
	"skufinder/internal"
)

func main() {
	fnder, err := internal.Init("https://images.asos-media.com/products/asos/204580792-2", &internal.ConfigNike)

	if err != nil {
		panic(err)
	}

	words, err := fnder.GetSku()

	if err != nil {
		panic(err)
	}

	fmt.Println(words)
}
