# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY aboutme_backend/go.mod ./
RUN go mod download
COPY aboutme_backend/. .
# Build the binary with static linking for Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY ./aboutme/dist ./dist
# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
