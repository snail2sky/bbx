ARG IMAGE

FROM $IMAGE as builder

WORKDIR /data0/src/bbx
ENV CGO_ENABLED=0
COPY . .
RUN go build


FROM alpine as runner

WORKDIR /data0/apps/bbx
COPY --from=builder /data0/src/bbx/bbx /data0/apps/bbx/bbx
ENTRYPOINT ["/data0/apps/bbx/bbx"]
