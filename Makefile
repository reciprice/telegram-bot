# Makefile based on template
# https://sohlich.github.io/post/go_makefile/

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOPATH=$(shell pwd)
export GOPATH

BINARY_NAME=telegram-bot
    
all: deps test build

build:
	$(GOBUILD) -o $(BINARY_NAME) telegram 
test: 
	$(GOTEST) telegram
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v telegram
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/go-telegram-bot-api/telegram-bot-api
	$(GOGET) gopkg.in/tucnak/telebot.v2