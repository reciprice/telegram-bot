FROM golang:1.12
WORKDIR /go/src/telegram
COPY . .
ENV GOPATH=/go/src/telegram
RUN go get github.com/go-telegram-bot-api/telegram-bot-api
RUN go get gopkg.in/tucnak/telebot.v2
RUN go run telegram