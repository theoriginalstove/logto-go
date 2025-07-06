package logto

//go:generate go tool oapi-codegen -config oapi.config.yml ./logto-openapi-source.json

type LogtoClient struct {
	BaseURL string
}

type LogtoClientConfig struct {
	BaseURL string
}

func NewLogtoClient(config LogtoClientConfig) *LogtoClient {
	return &LogtoClient{
		BaseURL: config.BaseURL,
	}
}
