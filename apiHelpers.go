package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type statsData struct {
	Sessions int `json:"sessions"`
	Total    int `json:"total"`
}

type sessionData struct {
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Time      int    `json:"time"`
	Elapsed   int    `json:"elapsed"`
	Remaining int    `json:"remaining"`
	EndTime   string `json:"endTime"`
	Paused    bool   `json:"paused"`
	Completed bool   `json:"completed"`
	Goal      string `json:"goal"`
	Work      string `json:"work"`
	MessageTs string `json:"messageTs"`
}

func getJsonString(endpoint string, slackId string, apiKey string) (string, error) {
	urlStr := fmt.Sprintf("https://hackhour.hackclub.com%s/%s", endpoint, slackId)
	client := &http.Client{}
	reqData := url.Values{}
	r, err := http.NewRequest(http.MethodGet, urlStr, strings.NewReader(reqData.Encode()))
	if err != nil {
		return "", fmt.Errorf("Error generating request: %v", err)
	}
	r.Header.Add("Authorization", "Bearer "+apiKey)
	resp, err := client.Do(r)
	if err != nil {
		return "", fmt.Errorf("Error making the request: %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error response body: %v", err)
	}
	err = resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("Error closing response body: %v", err)
	}
	return string(body), nil
}

func getSessionData(slackId string, apiKey string) (sessionData, error) {
	jsonData := struct {
		Ok   bool        `json:"ok"`
		Data sessionData `json:"data"`
	}{}
	body, err := getJsonString("/api/session", slackId, apiKey)
	if err != nil {
		return sessionData{}, nil
	}
	err = json.Unmarshal([]byte(body), &jsonData)
	if err != nil {
		return sessionData{}, fmt.Errorf("Error extracting json from response: %v", err)
	}

	return jsonData.Data, nil
}

func getStatsData(slackId string, apiKey string) (statsData, error) {
	jsonData := struct {
		Ok   bool      `json:"ok"`
		Data statsData `json:"data"`
	}{}
	body, err := getJsonString("/api/stats", slackId, apiKey)
	if err != nil {
		return statsData{}, nil
	}
	err = json.Unmarshal([]byte(body), &jsonData)
	if err != nil {
		return statsData{}, fmt.Errorf("Error extracting json from response: %v", err)
	}

	return jsonData.Data, nil
}
