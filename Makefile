OS := $(shell uname)

TARGET := windows

ifeq ($(OS),Darwin)
	TARGET := darwin
endif

ifeq ($(OS),Linux)
	TARGET := linux
endif


VERSION_FILE := cmd/version.go
TAG := $(shell sed -rn 's/const version = "(.*)"/\1/p' $(VERSION_FILE))
OS_LIST := windows darwin linux
ARCH_LIST := 386 amd64 arm64


auto: $(TARGET)

test:
		@echo $(IMAGE) $(TAG) $(GOOS) $(GOARCH) $(GOPROXY)

build:
		@export CGO_ENABLED=0 go mod tidy && go build

darwin:
		@echo "build for darwin"
		@env GOOS=darwin $(MAKE) build

linux:
		@echo "build for linux"
		@env GOOS=linux $(MAKE) build

windows:
		@echo "build for windows"
		@env GOOS=windows $(MAKE) build

install:
		@echo "install bbx"
		@go install

release:
		@echo Generate releases for $(TAG)
		@RELEASE_DIR=releases/$(TAG) && mkdir -p $$RELEASE_DIR && \
		export CGO_ENABLED=0 ; \
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
		@echo "build $(IMAGE) $(GOARCH) $(GOPROXY) $(TAG)"
		docker build --no-cache -t snail2sky/bbx:$(TAG) \
			--build-arg IMAGE=$(IMAGE)  \
			.

clean:
		@echo "clean bbx bbx.exe program"
		@rm -f bbx bbx.exe


clean-all:
		@echo "clean all binary file, such bbx, bbx.exe and all releases"
		@rm -rf bbx bbx.exe releases
		@docker image rm snail2sky/bbx:$(TAG)

help:
		@echo "build:           make [ darwin | linux | windows ]   build for mac linux or windows"
		@echo "clean:           make clean                          clean bbx and bbx.exe file"
		@echo "clean-all        make clean-all                      clean all binary files"
		@echo "install:         make install                        install binary file to GOBIN dir"
		@echo "build-image:     make build-image IMAGE=golang:1.22  build docker image"
		@echo "release:         make release                        generate releases into releases/$(TAG)"