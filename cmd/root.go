package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	filters = []string{}
	limit   int
	timeout int
	logging bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tzpipe",
	Short: "communicates with minibus channels",
	Long:  `tzpipe communicates with a minibus server to send and receive messages `,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initENV)
	rootCmd.PersistentFlags().StringArrayVarP(&filters, "filter", "f", []string{}, "regular expression filter (RE2, specifically https://github.com/google/re2/wiki/Syntax)")
	rootCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 0, "number of lines before exiting, default 0 (don't exit)")
	rootCmd.PersistentFlags().BoolVarP(&logging, "log", "L", false, "log to stderr, useful when piping stout")
	rootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 0, "number of seconds before exiting, default 0 (don't exit)")
}

func initENV() {

}
