FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN apk update && apk add git
RUN CGO_ENABLED=0 go build -o ./bot

FROM golang:1.18-alpine
WORKDIR /app

# for Heroku release phase logging
RUN apk --no-cache add curl

RUN apk --no-cache add --upgrade bash

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY ./makemigrate.sh .
RUN chmod +x makemigrate.sh

COPY ./push_messages.sh .
RUN chmod +x push_messages.sh

COPY ./models/migrations ./models/migrations

COPY --from=builder /app/bot .

# download the secure tunnel startup script from the borealis/postgres-ssh buildpack
RUN apk add --no-cache autossh && \
  wget https://raw.githubusercontent.com/OldSneerJaw/heroku-buildpack-borealis-pg-ssh/main/profile.d/borealis-pg-init-ssh-tunnel.sh && \
  chmod u+x borealis-pg-init-ssh-tunnel.sh

# set the command that establishes the SSH tunnel and launches the app when the container starts
CMD ./borealis-pg-init-ssh-tunnel.sh && ./bot
