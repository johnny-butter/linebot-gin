FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN apk update && apk add git
RUN CGO_ENABLED=0 go build -o ./bot

FROM golang:1.18-alpine
WORKDIR /app
COPY --from=builder /app/bot .
CMD [ "./bot" ]
