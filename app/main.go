package main

import (
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	// tgClient := telagam.New(tgBotHost, mustToken())
}

func mustToken () string {
	token := flag.String("token-bot-token", "", "Telegram Bot Access Token")

	flag.Parse()

	if *token == "" {
		log.Fatal("Telegram Bot Token is not specified")
	}

	return *token
}