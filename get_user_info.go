package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		panic(err)
		return
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer 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"},
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
