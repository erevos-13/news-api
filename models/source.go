package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/erevos-13/news-api-go/utils"
)

type SourceParams struct {
	Category string
	Language string
	Country  string
}

type Source struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

type ResponseSource struct {
	Status       string   `json:"status"`
	TotalResults int      `json:"totalResults"`
	Source       []Source `json:"sources"`
}

func SourceNew() *SourceParams {
	return &SourceParams{}
}

func (th SourceParams) GetSources() (ResponseSource, error) {
	baseUrl := fmt.Sprintf("https://newsapi.org/%s/top-headlines/sources", utils.GetVersion())
	params := []string{}
	params = append(params, fmt.Sprintf("apiKey=%s", utils.GetApiKey()))
	if th.Country != "" {
		params = append(params, fmt.Sprintf("country=%s", th.Country))
	}
	if th.Category != "" {
		params = append(params, fmt.Sprintf("category=%s", th.Category))
	}
	reqUrl := fmt.Sprintf("%s?%s", baseUrl, strings.Join(params, "&"))

	resp, err := http.Get(reqUrl)
	if err != nil {
		return ResponseSource{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseSource{}, err
	}

	var response ResponseSource
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ResponseSource{}, err
	}

	return response, nil
}
