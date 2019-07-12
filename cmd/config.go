package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "This command displays the configuration.",
	Long:  `This command displays the configuration for the CLI to use.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Configuration: %s\n", viper.ConfigFileUsed())
		fmt.Printf("Twitter Account: %s\n", viper.GetString("account"))
		fmt.Printf("Export File: %s\n", viper.GetString("file"))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
