package cmd

import (
	"fmt"

	"github.com/koh/ollama-client/pkg/ollama"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	prompt string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate text completion",
	Long:  `Generate text completion using an Ollama model.`,
	Run: func(cmd *cobra.Command, args []string) {
		if prompt == "" {
			fmt.Println("Error: prompt is required")
			return
		}
		runGenerate()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "Prompt for generation (required)")
	generateCmd.Flags().StringVarP(&model, "model", "m", "", "Model to use for generation (uses default if not specified)")
	generateCmd.Flags().Float64VarP(&temperature, "temperature", "t", 0.7, "Temperature for generation")
	generateCmd.MarkFlagRequired("prompt")
}

func runGenerate() {
	client := ollama.NewClient(GetAPIURL())

	if model == "" {
		model = viper.GetString("default_model")
		if model == "" {
			fmt.Println("Error: No model specified. Use --model flag or set a default with 'olc config set model <model>'")
			return
		}
	}

	req := &ollama.GenerateRequest{
		Model:  model,
		Prompt: prompt,
		Stream: false,
		Options: &ollama.Options{
			Temperature: temperature,
		},
	}

	fmt.Println("Generating response...")
	resp, err := client.Generate(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("\nResponse:")
	fmt.Println(resp.Response)

	if resp.EvalCount > 0 && resp.EvalDuration > 0 {
		tokensPerSecond := float64(resp.EvalCount) / (float64(resp.EvalDuration) / 1e9)
		fmt.Printf("\n[Tokens: %d, Speed: %.1f tokens/s]\n", resp.EvalCount, tokensPerSecond)
	}
}