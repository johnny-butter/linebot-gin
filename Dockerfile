FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN apk update && apk add git
RUN CGO_ENABLED=0 go build -o ./bot

FROM golang:1.18-alpine
WORKDIR /app

# for Heroku release phase logging
RUN apk --no-cache add curl

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY ./makemigrate.sh .
RUN chmod +x makemigrate.sh

COPY ./push_messages.sh .
RUN chmod +x push_messages.sh

COPY ./models/migrations ./models/migrations

COPY --from=builder /app/bot .

CMD [ "./bot" ]
