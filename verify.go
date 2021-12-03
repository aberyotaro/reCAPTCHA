package reCAPTCHA

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	Url = "https://www.google.com/recaptcha/api/siteverify"
)

type Response struct {
	ChallengeTs *string
	ErrorCodes []string
	Hostname *string
	Success *bool
}

func Verify(token *string) (*Response, error) {
	res, err := http.PostForm(Url, url.Values{
		"secret": {"your_secret_key"},
		"response": {*token},
	})
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	} (res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := &Response{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}

	return r, nil
}
