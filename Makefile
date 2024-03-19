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

clean:
		@echo "clean bbx bbx.exe program"
		@rm -f bbx bbx.exe

help:
		@echo "build:   make [ mac | linux | windows ]"
		@echo "clean:   make clean"
		@echo "install: make install"