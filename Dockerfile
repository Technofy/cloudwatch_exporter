#### Docker build image
FROM        golang:alpine AS builder
RUN         apk update && apk add --no-cache git
WORKDIR     /build
COPY        . .
RUN         CGO_ENABLED=0 GOOS=linux go build -o /out/cloudwatch_exporter


#### Docker image
FROM        alpine:latest
LABEL       maintainer="Anthony Teisseire <anthony.teisseire@technofy.io>"

COPY        --from=builder /out/cloudwatch_exporter /bin/cloudwatch_exporter
COPY        config.yml /etc/cloudwatch_exporter/config.yml

RUN         apk update && \
            apk add ca-certificates && \
            update-ca-certificates

EXPOSE      9042
ENTRYPOINT  [ "/bin/cloudwatch_exporter", "-config.file=/etc/cloudwatch_exporter/config.yml" ]