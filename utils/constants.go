package utils

var api_key string
var version string

func SetApiKey(apiKey string) {
	api_key = apiKey
}

func GetApiKey() string {
	return api_key
}

func SetVersion(ver string) {
	version = ver
}

func GetVersion() string {
	return version
}
