FROM golang:1.22.2-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 8082

CMD ["./main"]
