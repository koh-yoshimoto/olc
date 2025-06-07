package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "olc",
	Short: "A CLI client for Ollama API",
	Long:  `OLC (Ollama Client) - A command-line interface client for interacting with Ollama API in private network environments.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ollama-client.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ollama-client")
	}

	viper.SetDefault("ip", "localhost")
	viper.SetDefault("port", "11434")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func GetAPIURL() string {
	ip := viper.GetString("ip")
	port := viper.GetString("port")
	if port == "" {
		port = "11434"
	}
	return fmt.Sprintf("http://%s:%s", ip, port)
}