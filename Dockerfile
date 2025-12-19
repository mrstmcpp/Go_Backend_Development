#build stage
FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/main.go


#run stage
FROM golang:alpine
WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 3000
CMD ["./server"]
