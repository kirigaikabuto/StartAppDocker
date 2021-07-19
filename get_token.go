package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Token struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type Repo struct {
	Id       int64  `json:"id"`
	CloneUrl string `json:"clone_url"`
	Name     string `json:"name"`
}

func main() {
	clientId := "d8bd2c153ccb947d8208"
	clientSecret := "a844db007f4d22024fa2f041bab52bb83406f75b"
	code := "76406101c441a49a8ced"
	url := "https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s"
	url = fmt.Sprintf(url, clientId, clientSecret, code)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
		return
	}
	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
		return
	}
	resJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
		return
	}
	accessKey := strings.Split(strings.Split(string(resJson), "&")[0], "=")[1]
	fmt.Println(accessKey)
	reposUrl := "https://api.github.com/user/repos"
	token := accessKey
	client = http.Client{}
	req, err = http.NewRequest("GET", reposUrl, nil)
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
	var results []Repo
	err = json.Unmarshal(dJ, &results)
	if err != nil {
		panic(err)
		return
	}
	for _, v := range results {
		if v.Name == "cicd-buzz" {
			_, err := git.PlainClone("", false, &git.CloneOptions{
				URL:      v.CloneUrl,
				Progress: os.Stdout,
			})
			if err != nil {
				panic(err)
				return
			}
		}
	}
}
