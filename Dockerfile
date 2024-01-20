FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o whatismyip .

FROM alpine:latest
WORKDIR /
COPY --from=builder /app/whatismyip /whatismyip
EXPOSE 8080
ENTRYPOINT [ "/whatismyip" ]
