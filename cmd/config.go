package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings",
	Long:  `Manage configuration settings for IP, PORT, and default model.`,
}

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set configuration values",
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		showConfig()
	},
}

var setIPCmd = &cobra.Command{
	Use:   "ip [address]",
	Short: "Set Ollama server IP address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setConfigValue("ip", args[0])
	},
}

var setPortCmd = &cobra.Command{
	Use:   "port [number]",
	Short: "Set Ollama server port",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setConfigValue("port", args[0])
	},
}

var setModelCmd = &cobra.Command{
	Use:   "model [name]",
	Short: "Set default model",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setConfigValue("default_model", args[0])
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configShowCmd)
	
	configSetCmd.AddCommand(setIPCmd)
	configSetCmd.AddCommand(setPortCmd)
	configSetCmd.AddCommand(setModelCmd)
}

func setConfigValue(key, value string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		return
	}

	configPath := fmt.Sprintf("%s/.olc.yaml", home)
	
	viper.Set(key, value)
	
	if err := viper.WriteConfigAs(configPath); err != nil {
		fmt.Printf("Error saving configuration: %v\n", err)
		return
	}

	displayKey := key
	if key == "default_model" {
		displayKey = "Default model"
	} else if key == "ip" {
		displayKey = "IP address"
	} else if key == "port" {
		displayKey = "Port"
	}

	fmt.Printf("%s set to: %s\n", displayKey, value)
}

func showConfig() {
	fmt.Println("Current configuration:")
	fmt.Printf("  IP address:     %s\n", viper.GetString("ip"))
	fmt.Printf("  Port:           %s\n", viper.GetString("port"))
	fmt.Printf("  Default model:  %s\n", viper.GetString("default_model"))
	
	if cfgFile := viper.ConfigFileUsed(); cfgFile != "" {
		fmt.Printf("\nConfig file: %s\n", cfgFile)
	}
}