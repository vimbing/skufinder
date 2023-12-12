# Photo sku finder

This repo was used to find shoe skus by it's photo. It uses google lens to find results from photo and extracts possible skus, using config regexp and additional check function, so it can be used for any kind of photos, not only shoes.  

# Disclaimer
Since I didn't need to really extract exact informations, and I could just use regexps etc. It doesn't really format google response to readable format. It leads to situation, where looking for words is harder since bot doesn't know if analyzed word is part of text which is readable for user, or is any kind of google stuff.

# Example

    package main

      
    
    import (
	    "fmt"
	    "skufinder/internal/finder"
    )
 
    const IMG =  "https://img01.ztat.net/article/spp-media-p1/a3f04278eaca490ab8c740820f2fdae5/ef662e36acc144b1b09dd5a442dfac26.jpg?imwidth=1800&filter=packshot"
 
    func  main() {
	    f := finder.Init(IMG, &finder.ConfigNike)
	    
	    result, err := f.GetSku()
	    
	    if err !=  nil {
		    panic(err)
	    }
	    
	    for _, word :=  range result {
			fmt.Printf("%s - %d occurrences\n", word.Word, word.Count)
		}
    }
   
 ### Output
    Getting photo bytes...
    Photo bytes successfully scraped...
    Uploading photo to google lens...
    Photo successfully uploaded to google lens!
    Scraping google lens results...
    Google lens results successfully scraped!
    fv0384-001 - 25 occurrences
# Configs
There are some predefined configs for popular shoe brands, but you can set your own config like so
       
    ConfigNewBalance = Config{
			    SkuRegexp: regexp.MustCompile(`.*\d.*[a-zA-Z].*|.*[a-zA-Z].*\d.*`),
			    MinimumLength: 5,
			    MaximumLength: 15,
			    // Func check currently analyzed word and returns if it should be count
			    // or not.
			    AdditionalCheckFunc: func(s string) bool {
				    wordRegexp := regexp.MustCompile(`[a-zA-Z]`)
				    digitRegexp := regexp.MustCompile(`\d`)
				    return wordRegexp.MatchString(s) && digitRegexp.MatchString(s)
			    },
    }
