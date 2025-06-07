package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/koh-yoshimoto/olc/pkg/ollama"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	model       string
	temperature float64
	system      string
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start an interactive chat session",
	Long:  `Start an interactive chat session with an Ollama model.`,
	Run: func(cmd *cobra.Command, args []string) {
		runChat()
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	chatCmd.Flags().StringVarP(&model, "model", "m", "", "Model to use for chat (uses default if not specified)")
	chatCmd.Flags().Float64VarP(&temperature, "temperature", "t", 0.7, "Temperature for generation")
	chatCmd.Flags().StringVarP(&system, "system", "s", "", "System prompt")
}

func runChat() {
	client := ollama.NewClient(GetAPIURL())
	
	if model == "" {
		model = viper.GetString("default_model")
		if model == "" {
			fmt.Println("Error: No model specified. Use --model flag or set a default with 'olc config set model <model>'")
			return
		}
	}
	
	messages := []ollama.Message{}
	
	if system != "" {
		messages = append(messages, ollama.Message{
			Role:    "system",
			Content: system,
		})
	}

	reader := bufio.NewReader(os.Stdin)
	
	fmt.Printf("Chatting with %s. Type 'exit' to quit.\n", model)
	fmt.Println(strings.Repeat("-", 50))

	for {
		fmt.Print("\nYou: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		messages = append(messages, ollama.Message{
			Role:    "user",
			Content: input,
		})

		req := &ollama.ChatRequest{
			Model:    model,
			Messages: messages,
			Stream:   false,
			Options: &ollama.Options{
				Temperature: temperature,
			},
		}

		fmt.Print("\nAssistant: ")
		resp, err := client.Chat(req)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Println(resp.Message.Content)
		
		messages = append(messages, resp.Message)

		if resp.EvalCount > 0 && resp.EvalDuration > 0 {
			tokensPerSecond := float64(resp.EvalCount) / (float64(resp.EvalDuration) / 1e9)
			fmt.Printf("\n[Tokens: %d, Speed: %.1f tokens/s]\n", resp.EvalCount, tokensPerSecond)
		}
	}
}