OS := $(shell uname)
ifeq ($(OS),Windows_NT)
	TARGET := windows
else
	ifeq ($(OS),Darwin)
		TARGET := mac
	else
		TARGET := linux
	endif
endif

VERSION_FILE := cmd/version.go
TAG := $(shell sed -rn 's/const version = "(.*)"/\1/p' $(VERSION_FILE))
OS_LIST := windows darwin linux
ARCH_LIST := 386 amd64 arm64

auto: $(TARGET)

test:
		@echo $(TAG)

tidy:
		@go mod tidy

build:
		@go build


compile:
		@make tidy && make build

mac:
		@echo "build for mac"
		@export GOOS=darwin && make compile

linux:
		@echo "build for linux"
		@export GOOS=linux && make compile

windows:
		@echo "build for windows"
		@export GOOS=windows && make compile

install:
		@echo "install bbx"
		@go install

gen_releases:
		@echo Generate releases for $(TAG)
		@RELEASE_DIR=releases/$(TAG) && mkdir -p $$RELEASE_DIR && \
		for os in $(OS_LIST); do \
            for arch in $(ARCH_LIST); do \
              	exe=bbx-$$os-$$arch; \
                export GOOS=$$os GOARCH=$$arch; \
                if [ "$$os" = "windows" ]; then \
                  	echo "$$exe".exe; \
                    go build -o $$RELEASE_DIR/$$exe.exe; \
                else \
                  	echo "$$exe"; \
                    go build -o $$RELEASE_DIR/$$exe; \
                fi; \
            done ; \
        done


clean:
		@echo "clean bbx bbx.exe program"
		@rm -f bbx bbx.exe


clean_all:
		@echo "clean all binary file, such bbx, bbx.exe and all releases"
		@rm -rf bbx bbx.exe releases

help:
		@echo "build:        make [ mac | linux | windows ]  build for mac linux or windows"
		@echo "clean:        make clean                      clean bbx and bbx.exe file"
		@echo "clean_all     make clean_all                  clean all binary files"
		@echo "install:      make install                    install binary file to GOBIN dir"
		@echo "gen_releases: make gen_releases               generate releases into releases/$(TAG)"