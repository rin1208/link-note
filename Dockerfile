FROM golang:1.14


RUN go get github.com/oxequa/realize

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on
EXPOSE 8080


CMD ["go", "run", "backend/cmd/local.go"]