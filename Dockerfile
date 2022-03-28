FROM golang:1.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o ./bot

FROM golang:1.18
WORKDIR /app
COPY --from=builder /app/bot .
CMD [ "./bot" ]
