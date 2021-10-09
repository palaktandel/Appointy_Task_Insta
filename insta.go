package insta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var _ = fmt.Printf

type Config struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
	GrantType    string
	ResponseType string
	AccessToken  string

	// domain
	Domain string
	Prefix string
}

// client
type Instagram struct {
	Config *Config
	Client *http.Client
	Users  *UserApi
	Posts  *PostsApi
}

func (i *Instagram) SetAccessToken(accessToken string) {
	i.Config.AccessToken = accessToken
}

func (i *Instagram) NewRequest(item interface{}, method string, path string, params url.Values, isAccessToken bool) (*Content, error) {

	path = i.Config.Domain + i.Config.Prefix + path

	// create post parameter
	var bufferedBody *bytes.Buffer
	if strings.ToUpper(method) == "POST" {
		bufferedBody = bytes.NewBufferString(params.Encode())
		params = nil
	}

	if params == nil {
		params = url.Values{}
	}

	if isAccessToken {
		params.Set("access_token", i.Config.AccessToken)
	} else {
		params.Set("client_id", i.Config.ClientId)
	}

	var req *http.Request
	var err error
	if bufferedBody != nil {
		req, err = http.NewRequest(method, path, bufferedBody)
	} else {
		req, err = http.NewRequest(method, path, nil)
	}

	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = params.Encode()
	//fmt.Printf("%s", req.URL.String())

	if strings.ToUpper(method) == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := i.Client.Do(req)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("%v", resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("%s", body)

	var content = &Content{Data: item}
	json.Unmarshal(body, content)
	return content, nil
}

// make instance
func NewClient(callback func(*Config)) *Instagram {
	var config = new(Config)
	config.ResponseType = "code"
	config.GrantType = "authorization_code"
	config.Domain = "https://api.instagram.com"
	config.Prefix = "/v1"
	callback(config)

	var instagram = &Instagram{Config: config}
	instagram.Client = http.DefaultClient
	instagram.Users = &UserApi{Instagram: instagram}
	instagram.Posts = &PostsApi{Instagram: instagram}
	return instagram
}
