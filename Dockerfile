FROM golang:1.22-rc
WORKDIR /app
COPY . .
RUN go build -o bin ./cmd
CMD ["./bin"]