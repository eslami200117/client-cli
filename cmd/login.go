/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (

	"github.com/eslami200117/clientCli/handler"
	"github.com/spf13/cobra"
)

var (
	username string
	password string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login with username and password",
	Long: `login with username and password eg.
	clientCli login -u admin -p password`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.LoginHandler(username, password)

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "username for login")
	loginCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password for login")
	loginCmd.MarkPersistentFlagRequired("username")
	loginCmd.MarkPersistentFlagRequired("password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
