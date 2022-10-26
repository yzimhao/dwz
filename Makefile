
TARGET=dwz


build: linux_amd64
.PHONY: build

all: linux_amd64 windows_amd64 darwin_amd64 darwin_arm64

	

.PHONY: windows_amd64
windows_amd64:
	cd cmd/ && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../dist/$(TARGET).exe main.go

.PHONY: linux_amd64
linux_amd64:
	cd cmd/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../dist/$(TARGET) main.go

.PHONY: darwin_amd64
darwin_amd64:
	cd cmd/ && CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ../dist/$(TARGET) main.go

.PHONY: darwin_arm64
darwin_arm64:
	cd cmd/ && CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ../dist/$(TARGET) main.go 	