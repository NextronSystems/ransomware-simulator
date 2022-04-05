quickbuck.exe: $(wildcard cmd/* lib/*/*)
quickbuck.exe: export GOOS=windows
quickbuck.exe: export GOARCH=amd64

quickbuck.exe:
	go build -o $@ ./cmd

