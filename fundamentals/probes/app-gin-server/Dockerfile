FROM golang:1.23.5-bullseye AS builder
WORKDIR /app
COPY app.go go.mod go.sum .
RUN go mod download
RUN CGO_ENABLED=0 go build -o main

FROM scratch
EXPOSE 3000
COPY --from=builder /app/main /go/main
CMD ["/go/main"]
