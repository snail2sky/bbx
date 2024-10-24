ARG IMAGE=golang:1.22

FROM $IMAGE as builder

WORKDIR /data0/src/bbx

ENV CGO_ENABLED=0

COPY . .

# 静态编译
# RUN go build -ldflags '-linkmode "external" -extldflags "-static"'
RUN go build

FROM alpine as runner

WORKDIR /data0/apps/bbx

COPY --from=builder /data0/src/bbx/bbx /data0/apps/bbx/

ENTRYPOINT ["/data0/apps/bbx/bbx"]
