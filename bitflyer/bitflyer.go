package bitflyer

import (
	"net/http"
	"encoding/json"
	"bytes"
	"net/url"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"strconv"
	"io/ioutil"
	"time"
)

const baseURL = "https://api.bitflyer.com/v1/"

type APIClient struct {
	key string
	secret string
	httpClient *http.Client
}

func New(key, secret string) *APIClient {
	apiClient := &APIClient{key, secret, &http.Client{}}
	return apiClient
}

func(api APIClient) header(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println(timestamp)
	message := timestamp + method + endpoint + string(body)

	mac := hmac.New(sha256.New, []byte(api.secret))
	mac.Write([]byte(message))
	sign := hex.EncodeToString(mac.Sum(nil))
	return map[string]string {	
		"ACCESS-KEy": api.key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN": sign,
		"Content-Type": "application/json",
	}
}

func (api *APIClient) doRequest(method, urlPath string, query map[string]string, data[]byte) (body []byte, err error) {
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	endpoint := baseURL.ResolveReference(apiURL).String()
	log.Printf("action=doRequest endpoint=%s", endpoint)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	for key, value := range api.header(method, req.URL.RequestURI(), data) {
		req.Header.Add(key, value)
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type Balance struct {
	Currentcode string `json:"currency_code"`
	Amount float64 `json:"amount"`
	Available float64 `json:"available"`
}

func (api *APIClient) GetBalance() ([]Balance, error) {
	url := "me/getbalance"
	resp, err := api.doRequest("GET", url, map[string]string{}, nil)
	log.Printf("url=%s resp=%S", url, string(resp))
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	var balance []Balance
	err = json.Unmarshal(resp, &balance)
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	return balance, nil
}