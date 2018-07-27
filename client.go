package goflip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	imgFlipPostURL = "https://api.imgflip.com/caption_image"
)

// Client creates and sends requests to the ImgFlip API
type Client struct {
	username string
	password string
}

// NewClient create a new ImgFlip API client
func NewClient(username, password string) *Client {

	return &Client{username: username, password: password}
}

// SendRequest tries to send an ImgFlip request to the ImgFlip API and parse the response body
func (me *Client) SendRequest(request *APIRequest) (*APIResponse, error) {

	form := url.Values{}

	form.Add("template_id", request.TemplateID)
	form.Add("username", me.username)
	form.Add("password", me.password)
	form.Add("text0", "")
	form.Add("text1", "")

	for i, box := range request.Boxes {

		form.Add(fmt.Sprintf("boxes[%d][text]", i), box.Text)
		form.Add(fmt.Sprintf("boxes[%d][x]", i), fmt.Sprintf("%d", box.X))
		form.Add(fmt.Sprintf("boxes[%d][y]", i), fmt.Sprintf("%d", box.Y))
		form.Add(fmt.Sprintf("boxes[%d][width]", i), fmt.Sprintf("%d", box.Width))
		form.Add(fmt.Sprintf("boxes[%d][height]", i), fmt.Sprintf("%d", box.Height))
		form.Add(fmt.Sprintf("boxes[%d][color]", i), box.Color)
		form.Add(fmt.Sprintf("boxes[%d][outline_color]", i), box.OutlineColor)
	}

	req, err := http.NewRequest("POST", imgFlipPostURL, strings.NewReader(form.Encode()))

	if err != nil {
		return nil, fmt.Errorf("Error creating request: %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Error executing request: %s", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %s", err.Error())
	}

	response := &APIResponse{}

	err = json.Unmarshal(body, response)

	if err != nil {
		return nil, fmt.Errorf("Error parsing response body: %s", err.Error())
	}

	return response, nil
}
