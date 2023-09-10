package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"golang-first-project/lib/e"
	"io"
	"time"
)

type Storage interface {
	Save(page *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(page *Page) error
	IsExists(page *Page) (bool, error)
}

var ErrorsNoSavedPages = errors.New("no saved pages")

type Page struct {
	URL       string
	UserName  string
	CreatedAt time.Time
}

func (page Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, page.URL); err != nil {
		return "", e.Wrap("can't make hash from page's URL", err)
	}

	if _, err := io.WriteString(h, page.UserName); err != nil {
		return "", e.Wrap("can't make hash from page's UserName", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
