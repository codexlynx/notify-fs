FROM golang:1.18.3-buster AS builder
WORKDIR /go/src/github.com/codexlynx/notify-fs

RUN apt-get update -y \
    && apt-get install gcc-arm-linux-gnueabi -y \
    && mkdir -p /build/dist/

COPY . .

# amd64 build
RUN go build -o /build/dist/notify-fs_amd64 ./cmd/notify-fs

# arm build
RUN GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc go build -o /build/dist/notify-fs_arm ./cmd/notify-fs

# windows build
RUN GOOS=windows go build -o /build/dist/notify-fs_amd64.exe ./cmd/notify-fs

FROM scratch AS binary
COPY --from=builder /build/dist /
