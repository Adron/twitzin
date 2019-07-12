package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var followersCmd = &cobra.Command{
	Use:   "followers",
	Short: "This command displays the followers for the Twitter Account.",
	Long:  `This command displays the followers for the Twitter Account.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This needs implemented.")
	},
}

func init() {
	rootCmd.AddCommand(followersCmd)
}
