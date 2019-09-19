FROM golang:1.13 as builder

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go build -o /bin/main 


##############################################

FROM alpine:latest

RUN adduser -S app
USER app

COPY --from=builder /bin/main .

CMD [ "main" ]
