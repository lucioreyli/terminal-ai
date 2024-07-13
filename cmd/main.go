package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"terminal-ai/config"
	"terminal-ai/types"
)

func main() {
	req := types.Request{
		Model: config.Configs.Model,
		Messages: []types.Message{
			{
				Role:    "user",
				Content: []types.Content{{Type: "text", Text: config.Configs.Prompt}},
			},
		},
	}

	b, _ := json.Marshal(req)
	payload := bytes.NewBuffer(b)
	fmt.Println(string(payload.String()))
	request, err := http.NewRequest("POST", config.Envs.GPT_URL, payload)
	if err != nil {
		log.Fatal(err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+config.Envs.OPENAI_API_KEY)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(body))
}
