ARG IMAGE=golang:1.22

FROM $IMAGE as builder

ARG GOPROXY=goproxy.io
ARG GOOS=linux
ARG GOARCH=amd64

WORKDIR /data0/src/bbx

COPY . .

ENV GOPROXY=$GOPROXY GOOS=$GOOS GOARCH=$GOARCH

RUN echo "build for $GOOS $GOARCH"

# 静态编译
RUN go build -ldflags '-linkmode "external" -extldflags "-static"'


FROM alpine as runner

WORKDIR /data0/apps/bbx

COPY --from=builder /data0/src/bbx/bbx /data0/apps/bbx/

ENTRYPOINT ["/data0/apps/bbx/bbx"]