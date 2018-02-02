package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var netTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout: 5 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 5 * time.Second,
}

var netClient = &http.Client{
	Timeout:   time.Second * 10,
	Transport: netTransport,
}

type Data struct {
	Info Info `json:"data"`
}
type Api struct {
	server string
	port   int
	client http.Client
}

type Info struct {
	Version string `json:"version"`
}

func (a *Api) URL() string {
	return fmt.Sprintf("http://%s:%d", a.server, a.port)
}

func (a *Api) GetVersion() string {
	data := Data{}
	a.Request("GET", "", nil, &data)
	return data.Info.Version
}

func (a *Api) Request(method string, path string, data io.Reader, result interface{}) error {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", a.URL(), path), data)
	if err != nil {
		return err
	}
	resp, err := netClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return err
	} else {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		jsonErr := json.Unmarshal(body, &result)
		if jsonErr != nil {
			return jsonErr
		}
	}
	return nil
}
