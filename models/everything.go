package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/erevos-13/newsapigo/utils"
)

type Article struct {
	Source      struct{} `json:"source"`
	Author      string   `json:"author"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	URLToImage  string   `json:"urlToImage"`
	PublishedAt string   `json:"publishedAt"`
	Content     string   `json:"content"`
}

type Everything struct {
	SearchIn       string
	Domains        string
	From           string
	To             string
	Language       string
	SortBy         string
	ExcludeDomains string
	QueryParams
}

func EverythingNew() *Everything {
	return &Everything{}
}

func (ev Everything) GetEverything() (Response, error) {
	baseUrl := fmt.Sprintf("https://newsapi.org/%s/everything", utils.GetVersion())
	params := []string{}

	params = append(params, fmt.Sprintf("apiKey=%s", utils.GetApiKey()))
	params = append(params, fmt.Sprintf("q=%s", ev.KeywordsQuery))
	if ev.From != "" {
		params = append(params, fmt.Sprintf("from=%s", ev.From))
	}
	if ev.To != "" {
		params = append(params, fmt.Sprintf("to=%s", ev.To))
	}
	if ev.Language != "" {
		params = append(params, fmt.Sprintf("language=%s", ev.Language))
	}
	if ev.PageSize != 0 {
		params = append(params, fmt.Sprintf("pageSize=%d", ev.PageSize))
	}
	if ev.Page != 0 {
		params = append(params, fmt.Sprintf("page=%d", ev.Page))
	}
	if ev.SortBy != "" {
		params = append(params, fmt.Sprintf("sortBy=%s", ev.SortBy))
	}
	if ev.Domains != "" {
		params = append(params, fmt.Sprintf("domains=%s", ev.Domains))
	}
	finalUrl := fmt.Sprintf("%s?%s", baseUrl, strings.Join(params, "&"))
	res, err := http.Get(finalUrl)
	if err != nil {
		return Response{}, err
	}
	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}
	var responseJson Response
	err = json.Unmarshal([]byte(responseData), &responseJson)
	if err != nil {
		return Response{}, err
	}
	return responseJson, nil
}
