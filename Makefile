ransomware.exe: $(wildcard cmd/* lib/*/*)
ransomware.exe: GOOS=windows

ransomware.exe:
	go build -o ransomware.exe ./cmd

