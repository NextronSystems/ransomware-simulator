package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"ransomware-simulator/lib/encrypt"
	"ransomware-simulator/lib/note"
	"ransomware-simulator/lib/shadowcopy"
	"ransomware-simulator/lib/simulatemacro"

	"github.com/secDre4mer/go-parseflags"
	"github.com/spf13/cobra"
)

func init() {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run Ransomware Simulator",
		Run:   run,
	}
	runCmd.Flags().AddFlagSet(parseflags.CreateFlagset(&runOptions))
	rootCmd.AddCommand(runCmd)
}

var runOptions = struct {
	DisableMacroSimulation bool `flag:"disable-macro-simulation" description:"Don't simulate start from a macro by building the following process chain: winword.exe -> cmd.exe -> ransomware-simulator.exe"`

	DisableShadowCopyDeletion bool `flag:"disable-shadow-copy-deletion" description:"Don't simulate volume shadow copy deletion"`

	DisableFileEncryption bool   `flag:"disable-file-encryption" description:"Don't simulate document encryption"`
	EncryptionDirectory   string `flag:"dir" description:"Directory where files that will be encrypted should be staged"`

	DisableNoteDrop bool   `flag:"disable-note-drop" description:"Don't drop pseudo ransomware note"`
	NoteLocation    string `flag:"note-location" description:"Ransomware note location"`
}{
	EncryptionDirectory: `./encrypted-files`,
	NoteLocation:        filepath.Join(homeDir, "Desktop", "ransomware-simulator-note.txt"),
}

var homeDir, _ = os.UserHomeDir()

func run(cmd *cobra.Command, args []string) {
	if !runOptions.DisableMacroSimulation {
		// Simulate Macro execution of this executable with current parameters, including --disable-macro-simulation
		if err := simulatemacro.Run(append(os.Args, "--disable-macro-simulation")); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(asciiArt)
	if !runOptions.DisableShadowCopyDeletion {
		if err := shadowcopy.Delete(); err != nil {
			log.Fatal(err)
		}
	}
	if !runOptions.DisableFileEncryption {
		if err := encrypt.StageFiles(runOptions.EncryptionDirectory); err != nil {
			log.Fatal(err)
		}
		if err := encrypt.EncryptFiles(runOptions.EncryptionDirectory); err != nil {
			log.Fatal(err)
		}
	}
	if !runOptions.DisableNoteDrop {
		if err := note.Write(runOptions.NoteLocation); err != nil {
			log.Fatal(err)
		}
	}
}
