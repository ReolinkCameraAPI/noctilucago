FROM golang:1.14-alpine AS builder

RUN apk -U --no-cache add build-base git gcc

WORKDIR /build

ADD . .

RUN go mod download

RUN CGO_ENABLED=1 go build -a -o /usr/bin/noctiluca

FROM alpine:3.11

RUN addgroup -S noctiluca; \
    adduser -S noctiluca -G noctiluca -D  -h /home/noctiluca -s /bin/nologin; \
    chown -R noctiluca:noctiluca /home/noctiluca

WORKDIR /home/noctiluca

COPY --from=builder /usr/bin/noctiluca /usr/bin/noctiluca

VOLUME /home/noctiluca

EXPOSE 8000

USER noctiluca

ENTRYPOINT ["noctiluca"]