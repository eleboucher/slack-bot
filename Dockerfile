FROM golang:1.13 as builder

WORKDIR /go/src/github.com/genesixx/slack-bot

COPY . .


RUN CGO_ENABLED=0 go build -o /bin/main 

##############################################

FROM alpine:latest

RUN adduser -S app
USER app

WORKDIR /app
COPY --from=builder /bin/main .


CMD [ "./main" ]
