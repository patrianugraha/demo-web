#build stage
FROM golang:1.19.13-alpine AS builder

WORKDIR /workspace
COPY . .

RUN apk add --no-cache git
RUN go mod download
RUN go build -o app

#final stage
FROM alpine:latest
ENV TZ=Asia/Jakarta
RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /workspace/app /app

EXPOSE 8888
ENTRYPOINT ["./app"]