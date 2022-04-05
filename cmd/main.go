package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "ransomware-simulator { run | cleanup }",
	Example: "ransomware-simulator run",
	Short:   "Ransomware simulator",
}

//go:embed ascii-art.txt
var asciiArt string

func main() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(asciiArt)
		log.Print(err)
		os.Exit(1)
	}
}
