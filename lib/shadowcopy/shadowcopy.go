package shadowcopy

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func Delete() error {
	log.Println("Executing command 'vssadmin delete shadows /for=norealvolume /all /quiet'")
	err := exec.Command("vssadmin", "delete", "shadows", "/for=norealvolume", "/all", "/quiet").Run()
	// it takes some time to delete the shadowcopies
	time.Sleep(5 * time.Second)
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		if exitErr.ExitCode() != 2 {
			return fmt.Errorf("command returned unusual status code %d - might have been blocked", exitErr.ExitCode())
		}
	} else {
		return fmt.Errorf("could not run command: %w", err)
	}
	return nil
}
