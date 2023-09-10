package main

import (
	"flag"
	telegramCleint "golang-first-project/clients/telegram"
	event_consumer "golang-first-project/consumer/event-consumer"
	"golang-first-project/events/telegram"
	"golang-first-project/storage/files"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "users_pages"
	batchSize   = 100
)

func main() {
	tgClient := telegramCleint.New(tgBotHost, mustToken())

	eventsProcessor := telegram.New(tgClient, files.New(storagePath))

	log.Print("service started")

	if err := event_consumer.New(eventsProcessor, eventsProcessor, batchSize).Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "Telegram Bot Access Token")

	flag.Parse()

	if *token == "" {
		log.Fatal("Telegram Bot Token is not specified")
	}

	return *token
}
