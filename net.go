package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func makeMikrusApiCall(endpoint string) ([]byte, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		err    error
	)

	client = &http.Client{}

	data := url.Values{}
	data.Set("srv", config.Ctx[config.CurrentCtx].ServerName)
	data.Set("key", config.Ctx[config.CurrentCtx].Token)

	req, err = http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.mikr.us/%s", endpoint),
		strings.NewReader(data.Encode()),
	)

	if err != nil {
		return []byte(""), err
	}

	req.Header.Set("User-Agent", "mikruscli/0.1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err = client.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()

	respond, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respond, nil
}
