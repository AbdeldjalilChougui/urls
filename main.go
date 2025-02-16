package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://pub.dev:3000/pkgs/?user=garp&type=request"

func main () {
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	for _, val := range qparams {
		fmt.Println(val)
	}


	partsOfUrl := &url.URL {
		Scheme: "https",
		Host: "pub.dev",
		Path: "/pkgs/",
		RawQuery: "user=garp&type=request",		
	}

	fmt.Println(partsOfUrl)
}