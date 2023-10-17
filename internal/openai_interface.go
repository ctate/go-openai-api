package internal

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type OpenAIClientInterface interface {
	CreateChatCompletionStream(context.Context, openai.ChatCompletionRequest) (*openai.ChatCompletionStream, error)
}

var _ OpenAIClientInterface = &openai.Client{}
