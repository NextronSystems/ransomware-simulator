package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	stageCmd := &cobra.Command{
		Use:                "stage",
		Short:              "Run ransomware simulator via cmd /c",
		Run:                stage,
		DisableFlagParsing: true,
		Hidden:             true,
	}
	rootCmd.AddCommand(stageCmd)
}

func stage(cmd *cobra.Command, args []string) {
	log.Printf("Executing command via shell: %s", strings.Join(args, " "))
	// it takes some time to startup Winword.exe
	time.Sleep(5 * time.Second)
	stagedCommand := exec.Command(shell, shellParam, strings.Join(args, " "))
	stagedCommand.Stdout = os.Stdout
	stagedCommand.Stderr = os.Stderr
	if err := stagedCommand.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.ExitCode())
		}
		log.Fatal("Could not run staged command:", err)
	}
}
