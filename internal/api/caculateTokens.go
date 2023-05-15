package api

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/pandodao/tokenizer-go"
	"github.com/sirupsen/logrus"
)

func calculateTokens(input string) int {
	return tokenizer.MustCalToken(input)
}

func calcuReqTokens(body map[string]any) int {
	messages, ok := body["messages"].([]any)
	if !ok {
		return 0
	}
	input := ""
	for _, message := range messages {
		m, ok := message.(map[string]any)
		if !ok {
			return 0
		}
		v, ok := m["content"].(string)
		input += v
	}
	tokens := calculateTokens(input)
	return tokens
}

type Choice struct {
	Index        int    `json:"index"`
	FinishReason string `json:"finish_reason"`
	Delta        struct {
		Content string `json:"content"`
	} `json:"delta"`
}

type Chunk struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   string   `json:"usage"`
}

func calcuStreamRspTokens(buf *bytes.Buffer) int {
	// TODO: 这里全部读取了返回流，可能会占用较大内存，可以按\n切分流，逐步计算，直到最后的[DONE]
	bts, _ := io.ReadAll(buf)

	var contents string
	for _, line := range strings.Split(string(bts), "\n") {
		line = strings.Trim(line, " \t\n")
		if line != "" && strings.HasPrefix(line, "data") {
			line = line[6:]
			if line != "[DONE]" {
				var chunk Chunk
				if err := json.Unmarshal([]byte(line), &chunk); err != nil {
					logrus.Errorf("calcuRspTokens error:%v", err)
					return 0
				}

				for _, c := range chunk.Choices {
					contents += c.Delta.Content
				}
			}
		}
	}

	// logrus.Infof("rsp origin:%v", contents)
	tokens := calculateTokens(contents)
	return tokens
}
