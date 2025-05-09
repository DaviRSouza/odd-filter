FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -o oddfilter_app .

FROM scratch
WORKDIR /
COPY --from=builder /app/oddfilter_app /oddfilter_app
ENTRYPOINT ["/oddfilter_app"]