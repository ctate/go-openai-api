package main

import (
	"github.com/ctate/go-openai-api/internal"
)

func main() {
	internal.InitConfig()

	port := internal.GetConfig(internal.API_PORT)
	router := internal.Router()
	router.Run(":" + port)
}
