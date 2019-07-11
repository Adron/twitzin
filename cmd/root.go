package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"twitzin/helpers"
)

var configFile string

var rootCmd = &cobra.Command{
	Use: "twitzin",
	Short: "This is the CLI App to retrieve Twitter followers, following, and related for a particular account.",
	Long: `This CLI App can be used to retrieve Twitter followers, following, and other information from a particular
		Twitter account by passing in that particular twitter account.

For example:

twitzin followers @Adron

This command will retrieve the followers of @Adron's account.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	configFile = "twitzin.yaml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("twitzin")
	helpers.Check(viper.BindEnv("api_key"))
	helpers.Check(viper.BindEnv("api_secret"))

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}
}