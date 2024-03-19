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


mac:
		@echo "build for mac"
		@export GOOS=darwin && go build

linux:
		@echo "build for linux"
		@export GOOS=linux && go build

windows:
		@echo "build for windows"
		@export GOOS=windows && go build

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