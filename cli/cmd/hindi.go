/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hindiCmd represents the hindi command
var hindiCmd = &cobra.Command{
	Use:   "hindi",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("english") {
			fmt.Println("हेलो")
		} else {
			fmt.Println("नमस्ते")
		}
	},
}

// English indicates whether we just need to write "hello" in hindi
var English bool

func init() {
	rootCmd.AddCommand(hindiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hindiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	hindiCmd.Flags().BoolP("english", "e", false, "Write hello in hindi")
}
