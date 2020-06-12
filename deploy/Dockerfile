FROM golang:1.14-buster AS builder
RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y make upx
WORKDIR /go/gitlab.com/zephinzer/dev
COPY . /go/gitlab.com/zephinzer/dev
RUN make deps
RUN make setup_build
RUN make build_production
RUN make compress

FROM gcr.io/distroless/base-debian10 AS executable
COPY --from=builder /go/gitlab.com/zephinzer/dev/bin/dev_linux_amd64 /bin/dev
ENTRYPOINT ["/bin/dev"]
