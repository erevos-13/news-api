package newsapi

import (
	"github.com/erevos-13/news-api-go/models"
	"github.com/erevos-13/news-api-go/utils"
)

const API_KEY string = "default"

type NewsApi struct {
	Everything   models.Everything
	TopHeadlines models.TopHeadlines
	Source       models.SourceParams
}

func NewApi(apiKey string, version string) NewsApi {
	utils.SetApiKey(apiKey)
	utils.SetVersion(version)
	return NewsApi{
		Everything:   *models.EverythingNew(),
		TopHeadlines: *models.TopHeadlinesNew(),
		Source:       *models.SourceNew(),
	}
}
