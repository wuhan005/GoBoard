FROM golang:1.12 as go-builder

WORKDIR /app

COPY . .
ARG GOPROXY=https://goproxy.io

RUN go mod download \
    && pwd && ls \
    && go install github.com/gobuffalo/packr/packr && \
    CGO_ENABLED=0 packr build -o GoBoard


FROM alpine:latest as prod
COPY --from=go-builder /app/GoBoard /app/GoBoard

WORKDIR /app

EXPOSE 12306
ENTRYPOINT ["/app/GoBoard"]