package newsapi

import (
	"github.com/erevos-13/newsapigo/models"
	"github.com/erevos-13/newsapigo/utils"
)

const API_KEY string = "default"

type NewsApiResponse struct {
	Everything   models.Everything
	TopHeadlines models.TopHeadlines
	Source       models.SourceParams
}

func NewsApi(apiKey string, version string) NewsApiResponse {
	utils.SetApiKey(apiKey)
	utils.SetVersion(version)
	return NewsApiResponse{
		Everything:   *models.EverythingNew(),
		TopHeadlines: *models.TopHeadlinesNew(),
		Source:       *models.SourceNew(),
	}
}
