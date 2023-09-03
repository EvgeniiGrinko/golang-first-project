package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	t := mustToken()

	fmt.Sprintln(t)
}

func mustToken () string {
	token := flag.String("token-bot-token", "", "Telegram Bot Access Token")

	flag.Parse()

	if *token == "" {
		log.Fatal("Telegram Bot Token is not specified")
	}

	return *token
}
