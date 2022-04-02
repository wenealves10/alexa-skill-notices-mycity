FROM golang:1.16.6
RUN GOCACHE=OFF
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]
