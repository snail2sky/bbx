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

TAG := "v1.0.1"
OS_LIST := windows darwin linux
ARCH_LIST := 386 amd64 arm64

auto: $(TARGET)

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
		@for os in $(OS_LIST); do \
            for arch in $(ARCH_LIST); do \
              	exe=bbx-$$os-$$arch-$(TAG); \
                export GOOS=$$os GOARCH=$$arch; \
                if [ "$$os" = "windows" ]; then \
                  	echo "$$exe".exe; \
                    go build -o releases/$$exe.exe; \
                else \
                  	echo "$$exe"; \
                    go build -o releases/$$exe; \
                fi; \
            done ; \
        done


clean:
		@echo "clean bbx bbx.exe program"
		@rm -f bbx bbx.exe

help:
		@echo "build:   make [ mac | linux | windows ]"
		@echo "clean:   make clean"
		@echo "install: make install"