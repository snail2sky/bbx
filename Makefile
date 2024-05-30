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

IMAGE := golang:1.22
GOOS := linux
GOARCH := amd64
GOPROXY := goproxy.io

auto: $(TARGET)

test:
		@echo $(IMAGE) $(TAG) $(GOOS) $(GOARCH) $(GOPROXY)

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

gen-releases:
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

build-image:
		@echo "build docker image"
		@echo "build $(TAG) $(GOOS)"
		@docker build -t snail2sky/bbx:$(TAG) \
			--build-arg GOOS=$(GOOS) \
			--build-arg GOARCH=$(GOARCH) \
			--build-arg GOPROXY=$(GOPROXY) \
			--build-arg BUILD_TYPE=$(BUILD_TYPE) \
			--build-arg IMAGE=$(IMAGE) .

push-image:
		@echo "push docker image"
		@docker push snail2sky/bbx:$(TAG)

clean:
		@echo "clean bbx bbx.exe program"
		@rm -f bbx bbx.exe


clean-all:
		@echo "clean all binary file, such bbx, bbx.exe and all releases"
		@rm -rf bbx bbx.exe releases

help:
		@echo "build:        make [ mac | linux | windows ]  build for mac linux or windows"
		@echo "clean:        make clean                      clean bbx and bbx.exe file"
		@echo "clean-all     make clean-all                  clean all binary files"
		@echo "install:      make install                    install binary file to GOBIN dir"
		@echo "build-image:  make build-image [GOOS=[linux|darwin|windows] | GOARCH=[amd64|arm64|...] | GOPROXY=goproxy.io | IMAGE=golang:1.22]"
		@echo "push-image:   make push-image                 push image to docker hub"
		@echo "gen-releases: make gen-releases               generate releases into releases/$(TAG)"