package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func main() {
	resp, err := sendOP("我想了解一下活動", "OP1")
	log.Panicf("error:%v", err)

	log.Printf("%v", resp)
}

func sendOP(content string, opType string) (string, error) {

	if len(strings.TrimSpace(content)) == 0 {
		return "", fmt.Errorf("content of sentence is empty")
	}

	switch opType {
	case "OP1":
		content = fmt.Sprintf(`
		You are given a conversation between a user and a customer service bot, as well as the user's previous intention. By referring to the conversation, please update the user's intention. Their intention must be one of the following:
		["Conversation", "Ask about events", "Ask about mouse", "Ask about keyboard", "Ask about headphones"]. If the intention is unclear or does not fit any one of the choices, default to "Conversation". Only output one of the five choices without any elaboration.
		**User's intention:%v
		**Conversation**:%v
		`, "Conversation", content)
	case "OP2-1":
		content = fmt.Sprintf(`
 		Today is 2023/10/13. You are a customer service bot for an exhibition. The user is asking about events within the exhibition. Please extract the user's query and output it as a JSON object. The keys of the JSON object are "name", "type", "date", "time", "place" and "guests". "name" must be one of ["英雄召集令", "超級英雄聯盟賽", "2人銅心：實況主接力挑戰"] or "null" if no name is explicitly provided. There are only two "type": "offline" and "match". "date" must be in YYYY/MM/DD format. "time" must be 24-hour format. If there are missing details, default to "null". Only output the JSON object without line breaks and do not elaborate.
		User: "%v"
		`, content)
	case "OP2-2":
		content = fmt.Sprintf(`
		Today is 2023/10/13. You are a customer service bot for an exhibition. The user is asking about information about gaming mouse sold in the exhibition. Please extract the user's query and output it as a JSON object. The keys of the JSON object are "brand" (string), "name" (string), "hasLOL" (boolean), "dpi" (integer), "wireless" (boolean), "rgb" (boolean), "price_min" (float) and "price_max" (float). "name" is the actual product name or serial number. "hasLOL" is whether the mouse has the League of Legends (LOL) co-brand. "wireless" is whether the mouse is wireless or not. "rgb" is whether the mouse has RGB setting or not. "price_min" and "price_max" determine the price range of the mouse. If the price is ambiguous, provide a 500 yuan range difference. If there are missing details, default to "null". Only output the JSON object without line breaks and do not elaborate.
		User: "%v"
		`, content)
	case "OP2-3":
		content = fmt.Sprintf(`
		Today is 2023/10/13. You are a customer service bot for an exhibition. The user is asking about information about gaming keyboards sold in the exhibition. Please extract the user's query and output it as a JSON object. The keys of the JSON object are "brand" (string), "name" (string), "hasLOL" (boolean), "switch" (string), "wireless" (boolean), "numpad" (boolean), "rgb" (boolean), "price_min" (float) and "price_max" (float). "name" is the actual product name or serial number. "hasLOL" is whether the keyboard has the League of Legends (LOL) co-brand. "switch" is the keyboard switch brand and/or type such as "Razer黃軸". "wireless" is whether the keyboard is wireless or not. "numpad" is whether the keyboard has a number pad (full-size) or not. "rgb" is whether the keyboard has RGB setting or not. "price_min" and "price_max" determine the price range of the keyboard. If the price is ambiguous, provide a 500 yuan range difference. If there are missing details, default to "null". Only output the JSON object without line breaks and do not elaborate.
		User: "%v"
		`, content)
	case "OP2-4":
		content = fmt.Sprintf(`
		Today is 2023/10/13. You are a customer service bot for an exhibition. The user is asking about information about gaming headphones sold in the exhibition. Please extract the user's query and output it as a JSON object. The keys of the JSON object are "brand" (string), "name" (string), "hasLOL" (boolean), "frequency" (integer), "impedance" (integer), "type" (string), "wireless" (boolean), "rgb" (boolean), "price_min" (float) and "price_max" (float). "name" is the actual product name or serial number. "hasLOL" is whether the headphone has the League of Legends (LOL) co-brand. "frequency" is the headphone frequency in Hz. "impedance" is the headphone impedance in ohms. "type" must be either "Over-ear" (耳罩式) or "In-ear" (入耳式). "rgb" is whether the headphones has RGB setting or not. "price_min" and "price_max" determine the price range of the headphones. If the price is ambiguous, provide a 500 yuan range difference. If there are missing details, default to "null". Only output the JSON object without line breaks and do not elaborate.
		User: "%v"
		`, content)
	}

	log.Println("問題:", content)
	config := openai.DefaultAzureConfig("893af294fda047cdae8e7ee52c422e95", "https://goteam-sweden.openai.azure.com/openai/deployments/gpt35turbo/chat/completions?api-version=2023-07-01-preview")

	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,

			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("CreateChatCompletion error: %v", err)
	}

	respContent := resp.Choices[0].Message.Content

	return respContent, nil

}
