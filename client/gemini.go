package client

import (
	"context"

	"os"

	loging "github.com/Samarth1015/expense/Loging"
	"google.golang.org/genai"
)

var Geminiclient *genai.Client
var err error

func GeminiInit() {
	ctx := context.Background()
	Geminiclient, err = genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("OPENAI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		loging.Logger.Error("error in intialising gemini")
	}
}
