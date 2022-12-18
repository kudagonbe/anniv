/*
Copyright © 2022 Hikaru Imamoto
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// フラグバインド用の変数
var tag string

var rootCmd = &cobra.Command{
	Use:   "anniv",
	Short: "'anniv' is a CLI tool to manage anniversaries",
	Long: `'anniv' is a CLI tool to manage anniversaries.

You can register anniversaries and 
refer to the anniversaries list from the CLI.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//フラグの値を変数にバインド
	rootCmd.PersistentFlags().StringVar(&tag, "tag", "", "Tags for classifying anniversaries")
}
