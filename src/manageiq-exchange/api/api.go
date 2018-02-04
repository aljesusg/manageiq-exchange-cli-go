package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	info "manageiq-exchange/models/info"
	meta "manageiq-exchange/models/metadata"
	user "manageiq-exchange/models/user"
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

type Api struct {
	Server string
	Port   int
	Client *http.Client
	Data   DataApi
}

type DataApi struct {
	Data interface{}   `json:"data"`
	Meta meta.Metadata `json:"meta"`
}

func (a *Api) Init(server string, port int) {
	a.Server = server
	a.Port = port
	a.Client = netClient
}

func (a *Api) URL() string {
	return fmt.Sprintf("http://%s:%d", a.Server, a.Port)
}

func (a *Api) GetInfo() info.Info {
	a.Request("GET", "", nil)
	var info info.Info
	info.Init(a.Data.Data.(map[string]interface{}))
	return info
}

func (a *Api) GetUsers(expand bool) user.UserCollection {
	var path string
	if path = "/v1/users"; expand {
		path = "/v1/users?expand=resources"
	}
	err := a.Request("GET", path, nil)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	var users user.UserCollection
	users.Init(a.Data.Data.([]interface{}))
	return users
}

func (a *Api) Request(method string, path string, data io.Reader) error {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", a.URL(), path), data)
	if err != nil {
		return err
	}
	resp, err := a.Client.Do(req)
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
		jsonErr := json.Unmarshal(body, &a.Data)
		if jsonErr != nil {
			return jsonErr
		}
	}
	return nil
}
