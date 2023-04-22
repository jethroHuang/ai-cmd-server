package main

import (
	"ai-cmd-server/helper"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	openai "github.com/sashabaranov/go-openai"
	"os"
	"regexp"
)

const spell = "你现在是一个 %s 命令行高手，当我给你提出要求时，你只需要告诉我一条命令。注意，我不需要你对命令行进行解释，也不需要你用代码块包裹命令，当你认为无法写出对应的命令时，告诉我“此路不通！”。那么我的要求是：%s"

func main() {
	if len(os.Args) < 2 {
		app := os.Args[0]
		fmt.Println("Usage:")
		fmt.Println(app + " <your openai apikey>")
		return
	}
	apikey := os.Args[1]
	e := echo.New()
	e.GET("/help", func(c echo.Context) error {
		content := c.QueryParam("content")
		system := c.QueryParam("sys")
		if content == "" || system == "" {
			return helper.ErrorJson(c, "Required parameters: content,sys")
		}

		fullSpell := fmt.Sprintf(spell, system, content)

		client := openai.NewClient(apikey)
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: fullSpell,
					},
				},
			},
		)

		if err != nil {
			return helper.ErrorJson(c, fmt.Sprintf("ChatCompletion error: %v", err))
		}

		respContent := resp.Choices[0].Message.Content
		// 判断是否包含此路不通的回复
		unusable, _ := regexp.MatchString("此路不通", respContent)
		if unusable {
			return helper.ErrorJson(c, "Sorry, I can't help you.")
		}

		return helper.OkJson(c, resp.Choices[0].Message.Content)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
