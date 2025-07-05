package client

import (
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

var OpenAIClient *openai.Client

func InitOpenAI() {
	fmt.Println("apiiii", os.Getenv("OPENAI_API_KEY"))
	OpenAIClient = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	// resp, err := client.CreateChatCompletion(
	// 	context.Background(),
	// 	openai.ChatCompletionRequest{
	// 		Model: openai.GPT3Dot5Turbo,
	// 		Messages: []openai.ChatCompletionMessage{
	// 			{
	// 				Role:    openai.ChatMessageRoleUser,
	// 				Content: "Hello!",
	// 			},
	// 		},
	// 	},
	// )

	// if err != nil {
	// 	fmt.Printf("ChatCompletion error: %v\n", err)
	// 	return
	// }

	// fmt.Println(resp.Choices[0].Message.Content)
}
