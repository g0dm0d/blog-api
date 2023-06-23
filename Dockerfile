### STAGE 1: Build ###
FROM golang:alpine AS builder
WORKDIR /build
ADD ./blog_server/go.mod ./
COPY ./blog_server ./
RUN go build -o nttt_blog ./cmd/main.go


### STAGE 2: Run ###
FROM alpine:3.18.2
COPY --from=builder /build/nttt_blog nttt_blog
COPY --from=builder /build/config.toml config.toml
RUN mkdir assets
CMD ["./nttt_blog"]
