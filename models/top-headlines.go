package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/erevos-13/newsapigo/utils"
)

type TopHeadlines struct {
	Country  string
	Category string
	QueryParams
}

func TopHeadlinesNew() *TopHeadlines {
	return &TopHeadlines{}
}

func (th TopHeadlines) GetTopHeadlines() (Response, error) {
	baseUrl := fmt.Sprintf("https://newsapi.org/%s/top-headlines", utils.GetVersion())
	params := []string{}

	params = append(params, fmt.Sprintf("apiKey=%s", utils.GetApiKey()))
	if th.KeywordsQuery != "" {
		params = append(params, fmt.Sprintf("q=%s", th.KeywordsQuery))
	}
	queryParams := map[string]string{
		"country":  th.Country,
		"category": th.Category,
	}
	for key, value := range queryParams {
		if value != "" {
			params = append(params, fmt.Sprintf("%s=%s", key, value))
		}
	}
	if th.PageSize != 0 {
		params = append(params, fmt.Sprintf("pageSize=%d", th.PageSize))
	}
	if th.Page != 0 {
		params = append(params, fmt.Sprintf("page=%d", th.Page))
	}
	reqUrl := fmt.Sprintf("%s?%s", baseUrl, strings.Join(params, "&"))

	resp, err := http.Get(reqUrl)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}
