/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/eslami200117/clientCli/app/handler"
	"github.com/spf13/cobra"
)

var addSource string
// addSourceCmd represents the addSource command
var addSourceCmd = &cobra.Command{
	Use:   "addSource",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.AddSource(username, addSource, password)
	},
}

func init() {
	rootCmd.AddCommand(addSourceCmd)
	addSourceCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "username for login")
	addSourceCmd.PersistentFlags().StringVarP(&addSource, "addSource", "s", "", "name  for user")
	addSourceCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password for login")
	addSourceCmd.MarkPersistentFlagRequired("username")
	addSourceCmd.MarkPersistentFlagRequired("addSource")
	addSourceCmd.MarkPersistentFlagRequired("password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addSourceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addSourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
