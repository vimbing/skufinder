package main

import (
	"fmt"
	"skufinder/internal/finder"
)

const IMG = "https://img01.ztat.net/article/spp-media-p1/a3f04278eaca490ab8c740820f2fdae5/ef662e36acc144b1b09dd5a442dfac26.jpg?imwidth=1800&filter=packshot"

func main() {
	f := finder.Init(IMG, &finder.ConfigNike)

	result, err := f.GetSku()

	if err != nil {
		panic(err)
	}

	for _, word := range result {
		fmt.Printf("%s - %d occurrences\n", word.Word, word.Count)
	}
}
