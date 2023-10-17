package internal

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

type ChatStreamRequestBody struct {
	MaxTokens int    `json:"maxTokens"`
	Model     string `json:"model"`
	Prompt    string `json:"prompt" binding:"required"`
}

func ChatStream(client *openai.Client, c *gin.Context) {
	var body ChatStreamRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.MaxTokens == 0 {
		body.MaxTokens = 20
	}
	if body.Model == "" {
		body.Model = openai.GPT3Dot5Turbo
	}

	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     body.Model,
		MaxTokens: body.MaxTokens,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: body.Prompt,
			},
		},
		Stream: true,
	}
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("ChatCompletionStream error: %v", err)})
		return
	}
	defer stream.Close()

	w := c.Writer
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Stream error: %v", err)))
			return
		}

		w.Write([]byte(response.Choices[0].Delta.Content))
		w.(http.Flusher).Flush()
	}
}
