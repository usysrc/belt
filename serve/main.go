package main

import (
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var port string

	rootCmd := &cobra.Command{
		Use:   "file-server",
		Short: "Start the server",
		Run: func(_ *cobra.Command, _ []string) {
			if err := startFileServer(port); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		},
	}
	rootCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port for the file server")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func startFileServer(port string) error {
	handler := http.FileServer(http.Dir("."))
	log.Printf("File server is running on http://localhost:%s\n", port)
	return http.ListenAndServe(":"+port, handler)
}
