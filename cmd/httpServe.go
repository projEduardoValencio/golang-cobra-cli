/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/projEduardoValencio/cli-golang-cobra/internal/webserver"
	"github.com/spf13/cobra"
)

var port string

// httpServeCmd represents the httpServe command
var httpServeCmd = &cobra.Command{
	Use:   "httpServe",
	Short: "Criação de um servidor HTTP.",
	Long:  `Criação de um servidor HTTP.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Atribuindo porta
		server := webserver.Server{Port: port}
		// Iniciando servidor
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpServeCmd)

	httpServeCmd.Flags().StringVarP(&port, "porta", "p", "4090", "Porta para o server, por padrão será 4090")
}
