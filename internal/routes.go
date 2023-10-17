package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	client := openai.NewClient(GetConfig(OPEN_API_KEY))
	router.POST("/chat", func(c *gin.Context) {
		ChatStream(client, c)
	})

	return router
}
