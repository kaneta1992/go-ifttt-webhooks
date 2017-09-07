package ifttt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const WEBHOOKS_URL = "https://maker.ifttt.com/trigger/%s/with/key/%s"

type Data struct {
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
	Value3 string `json:"value3"`
}

type Client struct {
	key string
}

func New(key string) *Client {
	return &Client{key: key}
}

func (self *Client) Post(event, value1, value2, value3 string) error {
	data := Data{value1, value2, value3}
	jsonStr, _ := json.Marshal(data)

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(WEBHOOKS_URL, event, self.key),
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

func (self *Client) PostWithBr(event, value1, value2, value3 string) error {
	return self.Post(event, crlfToBr(value1), crlfToBr(value2), crlfToBr(value3))
}

func crlfToBr(str string) string {
	tmp := strings.Replace(str, "\r\n", "<br>", -1)
	return strings.Replace(tmp, "\n", "<br>", -1)
}
