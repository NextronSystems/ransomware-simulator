package encrypt

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//go:embed TestDocument.docx
var document []byte

const documentCount = 10000

func StageFiles(directory string) error {
	log.Printf("Clearing target folder %s...", directory)
	if err := os.RemoveAll(directory); err != nil {
		return fmt.Errorf("could not delete folder: %w", err)
	}
	if err := os.MkdirAll(directory, 0755); err != nil {
		return fmt.Errorf("could not create folder: %w", err)
	}
	log.Printf("Writing %d documents to target folder %s...", documentCount, directory)
	for i := 0; i <= documentCount; i++ {
		if err := os.WriteFile(filepath.Join(directory, fmt.Sprintf("document%d.docx", i)), document, 0644); err != nil {
			return fmt.Errorf("could not write file: %w", err)
		}
	}
	return nil
}
