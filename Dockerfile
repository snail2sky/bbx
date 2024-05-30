ARG IMAGE=golang:1.22
ARG GOPROXY
ARG GOOS
ARG GOARCH

FROM $IMAGE as builder

WORKDIR /data0/src/bbx

COPY . .

ENV GOPROXY=${GOPROXY:-goproxy.io} GOOS=${GOOS:-linux} GOARCH=${GOARCH:-amd64}

RUN echo build for $GOOS $GOARCH

# 静态编译
RUN go build -ldflags '-linkmode "external" -extldflags "-static"'

FROM alpine as runner

WORKDIR /data0/apps/bbx

COPY --from=builder /data0/src/bbx/bbx /data0/apps/bbx/

ENTRYPOINT ["/data0/apps/bbx/bbx"]