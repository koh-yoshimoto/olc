package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/koh/ollama-client/pkg/ollama"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Manage Ollama models",
	Long:  `Manage Ollama models - list, pull, or delete models.`,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available models",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		listModels()
	},
}

var pullCmd = &cobra.Command{
	Use:   "pull [model]",
	Short: "Pull a model",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pullModel(args[0])
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [model]",
	Short: "Delete a model",
	Aliases: []string{"rm"},
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deleteModel(args[0])
	},
}

var setCmd = &cobra.Command{
	Use:   "set [model]",
	Short: "Set default model",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setDefaultModel(args[0])
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
	modelCmd.AddCommand(listCmd)
	modelCmd.AddCommand(pullCmd)
	modelCmd.AddCommand(deleteCmd)
	modelCmd.AddCommand(setCmd)
}

func listModels() {
	client := ollama.NewClient(GetAPIURL())

	resp, err := client.ListModels()
	if err != nil {
		fmt.Printf("Error listing models: %v\n", err)
		return
	}

	if len(resp.Models) == 0 {
		fmt.Println("No models found")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "NAME\tSIZE\tMODIFIED")
	
	for _, model := range resp.Models {
		size := formatBytes(model.Size)
		modified := formatTime(model.ModifiedAt)
		fmt.Fprintf(w, "%s\t%s\t%s\n", model.Name, size, modified)
	}
	
	w.Flush()
}

func pullModel(modelName string) {
	client := ollama.NewClient(GetAPIURL())

	fmt.Printf("Pulling model %s...\n", modelName)
	err := client.PullModel(modelName)
	if err != nil {
		fmt.Printf("Error pulling model: %v\n", err)
		return
	}

	fmt.Printf("Successfully pulled model %s\n", modelName)
}

func deleteModel(modelName string) {
	client := ollama.NewClient(GetAPIURL())

	fmt.Printf("Deleting model %s...\n", modelName)
	err := client.DeleteModel(modelName)
	if err != nil {
		fmt.Printf("Error deleting model: %v\n", err)
		return
	}

	fmt.Printf("Successfully deleted model %s\n", modelName)
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func formatTime(t time.Time) string {
	duration := time.Since(t)
	
	if duration < time.Minute {
		return "just now"
	} else if duration < time.Hour {
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	} else if duration < 24*time.Hour {
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	} else if duration < 30*24*time.Hour {
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	} else {
		return t.Format("2006-01-02")
	}
}

func setDefaultModel(modelName string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		return
	}

	configPath := fmt.Sprintf("%s/.ollama-client.yaml", home)
	
	viper.Set("default_model", modelName)
	
	if err := viper.WriteConfigAs(configPath); err != nil {
		fmt.Printf("Error saving default model: %v\n", err)
		return
	}

	fmt.Printf("Default model set to: %s\n", modelName)
}