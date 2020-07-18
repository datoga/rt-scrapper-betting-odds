package main

import "github.com/go-resty/resty/v2"

func getURL() ([]byte, error) {
	client := resty.New()

	jsonParam := `{"eventState":"Mixed","eventTypes":["Outright"],"ids":["152"],"eventTags":[],"leagueState":"Regular","pagination":{"top":100,"skip":0}}`

	resp, err := client.R().
		SetBody(jsonParam).
		SetHeader("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36").
		SetHeader("accept", "*/*").
		SetHeader("content-type", "application/json-patch+json").
		SetAuthScheme("Bearer").
		SetAuthToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTaXRlSWQiOjU2LCJTZXNzaW9uSWQiOiI4NzA5ZWNjMS01ZjM1LTQ4NDEtOThmZC1hZjBmYzNmZmI3ZDQiLCJuYmYiOjE1OTUwNTM2NjgsImV4cCI6MTU5NTY1ODQ5OCwiaWF0IjoxNTk1MDUzNjk4fQ.2yym9FNqjK7Lcua6SrVKj5J0QwZqla7DqFs-dvD7y-A").
		Post("https://sbapi.sbtech.com/10betcom/sportscontent/sportsbook/v1/Events/GetBySportId")

	if err != nil {
		return []byte{}, err
	}

	return resp.Body(), nil
}
