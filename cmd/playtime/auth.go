package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func auth(config AuthConfig) {
	fmt.Println(config.clientId)
	client := &http.Client{}
	data := url.Values{}
	data.Add("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, config.tokenUrl, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Could not create request object: ", err)
	}
	req.SetBasicAuth(config.clientId, config.clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request to tokenurl failed: ", err)
	}

	bodyText, _ := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	fmt.Println(s)
}
