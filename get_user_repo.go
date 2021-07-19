package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.github.com/user/repos"
	token := "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"
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
		return
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer " + token},
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
		return
	}
	dJ, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(dJ))
}
