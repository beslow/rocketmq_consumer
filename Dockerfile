FROM golang:1.19 as builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .    

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

RUN mkdir publish && \
    cp rocketmq_consumer publish && \
    cp config.yml publish && \
    mkdir publish/logs && \
    touch publish/logs/info.log

FROM alpine:3.14

RUN apk add wait4x

WORKDIR /app

COPY --from=builder /app/publish .

# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/cert

# ENV SOME_ENV=some_env \

# EXPOSE 80