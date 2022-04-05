package simulatemacro

import (
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Run(args []string) error {
	wd, _ := os.Getwd()
	winword := filepath.Join(wd, "WINWORD.EXE")
	// Create pseudo winword.exe - it's only a copy of this executable
	log.Println("Copying executable as pseudo WINWORD.EXE")
	if err := copyExecutable(winword); err != nil {
		return err
	}
	log.Printf("Staging execution via WINWORD.EXE: %s", strings.Join(args, " "))
	wordCommand := exec.Command(winword, append([]string{"stage"}, args...)...)
	wordCommand.Stdout = os.Stdout
	wordCommand.Stderr = os.Stderr
	if err := wordCommand.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.ExitCode())
		}
		log.Fatal("Could not run staged command:", err)
	} else {
		os.Exit(0)
	}
	return nil
}

func copyExecutable(target string) error {
	executable, err := os.Executable()
	if err != nil {
		return err
	}
	execFile, err := os.Open(executable)
	if err != nil {
		return err
	}
	defer execFile.Close()
	targetFile, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer targetFile.Close()
	_, err = io.Copy(targetFile, execFile)
	return err
}
