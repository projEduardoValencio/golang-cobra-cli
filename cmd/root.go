/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-golang-cobra",
	Short: "GO Lang com Cobra-CLI",
	Long:  `Este é uma aplicação de exemplo, apresentando uma construção simples com Cobra-CLI`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
