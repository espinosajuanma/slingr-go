package slingr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Environment string

const (
	EnvProd    = "prod"
	EnvDev     = "dev"
	EnvStating = "staging"
)

type App struct {
	Name       string
	Env        Environment
	Email      string
	Token      string
	httpClient *http.Client
}

func NewApp(name string, env Environment) *App {
	return &App{
		Name:       name,
		Env:        env,
		httpClient: &http.Client{},
	}
}

func (c *App) getURI(path string) string {
	return "https://" + c.Name + ".slingrs.io/" + string(c.Env) + "/runtime/api" + path
}

func (c *App) request(method, path string, payload interface{}, query map[string]string) ([]byte, error) {
	if c.Token == "" && path != "/auth/login" {
		return nil, fmt.Errorf("needs a token before making a request")
	}
	req, _ := http.NewRequest(method, c.getURI(path), nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if path != "/auth/login" {
		req.Header.Add("token", c.Token)
	}

	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	if payload != nil {
		data, _ := json.Marshal(payload)
		req.Body = ioutil.NopCloser(bytes.NewReader(data))
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		apiError = APIError{
			URL:        req.URL.String(),
			HttpMethod: req.Method,
			HttpStatus: res.StatusCode,
		}
		json.NewDecoder(res.Body).Decode(&apiError)
		return nil, apiError
	}

	responseData, _ := ioutil.ReadAll(res.Body)
	return responseData, nil
}

func (c *App) Get(path string, queryParams map[string]string) ([]byte, error) {
	return c.request("GET", path, nil, queryParams)
}

func (c *App) Post(path string, body interface{}, queryParams map[string]string) ([]byte, error) {
	return c.request("POST", path, body, queryParams)
}

func (c *App) Put(path string, body interface{}, queryParams map[string]string) ([]byte, error) {
	return c.request("PUT", path, body, queryParams)
}

func (c *App) Delete(path string, queryParams map[string]string) ([]byte, error) {
	return c.request("DELETE", path, nil, queryParams)
}
