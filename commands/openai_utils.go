package commands

import (
	"slices"
)

var apiURL = "https://api.openai.com/v1/"

var modelsMap = map[string][]string{
	"chat":      {"gpt-3.5-turbo-0125", "gpt-3.5-turbo", "gpt-4o"},
	"embedding": {"text-embedding-3-large", "text-embedding-3-small", "text-embedding-ada-002"},
}

func IsOpenAIModelAvailable(command string, modelName *string) bool {

	return slices.Contains(modelsMap[command], *modelName)

}
