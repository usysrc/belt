package main

import (
	"fmt"
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
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	rootCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port for the file server")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func startFileServer(port string) error {
	handler := http.FileServer(http.Dir("."))
	fmt.Printf("File server is running on http://localhost:%s\n", port)
	return http.ListenAndServe(":"+port, handler)
}
