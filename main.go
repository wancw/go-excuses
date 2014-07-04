package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.programmerexcuses.com/")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	re := regexp.MustCompile(`<a[^>]* href="/"[^>]*>(.+)</a>`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) != 2 {
		fmt.Println("No Excuses!")
		os.Exit(-1)
	}

	fmt.Println(matches[1])
}
