package main

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type ApiToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

type ApiTokenService interface {
	GetToken() ApiToken
}

type tokenService struct {
	cacheKey string
	config   AuthConfig
	cache    *cache.Cache
}

func NewApiTokenService(config AuthConfig) ApiTokenService {
	cache := cache.New(5*time.Minute, 10*time.Minute)
	return &tokenService{
		cacheKey: "apiToken",
		config:   config,
		cache:    cache,
	}
}

func (s *tokenService) GetToken() ApiToken {
	if token, found := s.cache.Get(s.cacheKey); found {
		return token.(ApiToken)
	} else {
		token := RetrieveApiToken(s.config)
		s.cache.Set(s.cacheKey, token, time.Duration(token.ExpiresIn)*time.Second)
		return token
	}
}

func RetrieveApiToken(config AuthConfig) ApiToken {
	client := &http.Client{}
	data := url.Values{}
	data.Add("grant_type", config.grantType)

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
	token := ApiToken{}

	errj := json.Unmarshal(bodyText, &token)
	if errj != nil {
		fmt.Println("Could not pase json to ApiToken struct")
	}

	return token
}
