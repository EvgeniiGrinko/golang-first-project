package telegram

import (
	"golang-first-project/clients/telegram"
	"golang-first-project/storage"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}
