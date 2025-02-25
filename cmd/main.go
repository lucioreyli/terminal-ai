package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"terminal-ai/color"
	"terminal-ai/config"
	"terminal-ai/types"
	"terminal-ai/utils"
)

func main() {
	if ok := utils.CheckValidModel(config.Configs.Model); !ok {
		fmt.Println(color.Format(color.RED, "Invalid model, check https://platform.openai.com/docs/models"))
		return
	}

	fmt.Println(fmt.Sprintf("%s:", color.Format(color.CYAN, "Model")), *config.Configs.Model)
	fmt.Println(fmt.Sprintf("%s:", color.Format(color.CYAN, "Prompt")), *config.Configs.Prompt)
	fmt.Println(fmt.Sprintf("%s:", color.Format(color.CYAN, "Temperature")), *config.Configs.Temperature)

	// source: https://platform.openai.com/docs/api-reference/chat
	req := types.Request{
		Model: *config.Configs.Model,
		Messages: []types.Message{
			{
				Role:    "user",
				Content: []types.Content{{Type: "text", Text: *config.Configs.Prompt}},
			},
		},
		Temperature: *config.Configs.Temperature,
	}

	b, err := json.Marshal(&req)
	if err != nil {
		fmt.Println(err)
		return
	}
	payload := bytes.NewBuffer(b)

	request, err := http.NewRequest("POST", config.Envs.GPT_URL, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+config.Envs.OPENAI_API_KEY)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	reader, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var res types.Response
	json.Unmarshal(reader, &res)

	fmt.Println(
		color.Format(color.YELLOW, res.Choices[0].Message.Content),
	)
}
