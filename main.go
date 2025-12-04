package main


import (
    "context"
    "fmt"
    "log"
    "google.golang.org/genai"
	// "github.com/openai/openai-go/v3"
	// "github.com/openai/openai-go/v3/option"
	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
)
func query_one_word(ctx context.Context, client *genai.Client, word string) string{

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text(word + " 这个词有哪些词性？每个词性对应的中文释义有哪些？纯文本形式返回结果，不要加markdown语法的字符"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
		return ""
    }
    return result.Text()
}
func main() {
	client, _ := deepseek.NewClient("My-api-key")

	chatReq := &request.ChatCompletionsRequest{
		Model:  deepseek.DEEPSEEK_CHAT_MODEL,
		Stream: false,
		Messages: []*request.Message{
			{
				Role:    "user",
				Content: "impose是什么意思?", // set your input message
			},
		},
	}

	chatResp, err := client.CallChatCompletionsChat(context.Background(), chatReq)
	if err != nil {
		fmt.Println("Error =>", err)
		return
	}
	fmt.Printf("%s\n", chatResp.Choices[0].Message.Content)
   	// ctx := context.Background()
    // // The client gets the API key from the environment variable `GEMINI_API_KEY`.
    // client, err := genai.NewClient(ctx, nil)
    // if err != nil {
    //     log.Fatal(err)
    // }
	// resp := query_one_word(ctx,client,"happy")
	// fmt.Println(resp)
	// client := openai.NewClient(
	// 	option.WithAPIKey("my-api-key"), // defaults to os.LookupEnv("OPENAI_API_KEY")
	// )
	// chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
	// 	Messages: []openai.ChatCompletionMessageParamUnion{
	// 		openai.UserMessage("今天吃啥"),
	// 	},
	// 	Model: openai.ChatModelGPT5,
	// })
	// if err != nil {
	// 	panic(err.Error())
	// }
	// println(chatCompletion.Choices[0].Message.Content)
}