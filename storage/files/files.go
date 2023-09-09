package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"golang-first-project/lib/e"
	"golang-first-project/storage"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return Storage{basePath: basePath}

}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.Wrap("can't save file", err) }()

	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.Mkdir(fPath, defaultPerm); err != nil {
		return err
	}
	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil

}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() { err = e.Wrap("can't save pick random page", err) }()
	fPath := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(fPath)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, storage.ErrorsNoSavedPages
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	file := files[n]

	return s.decodePage(filepath.Join(fPath, file.Name()))
}

func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, e.Wrap("can't open page", err)
	}

	defer func() { _ = f.Close }()

	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, e.Wrap("can't decode page", err)
	}

	return &p, nil

}
func (s Storage) Remove(page *storage.Page) error {
	fileName, err := fileName(page)
	if err != nil {
		return e.Wrap("can't get fileName of page to remove", err)
	}

	path := filepath.Join(s.basePath, page.UserName, fileName)

	if err := os.Remove(path); err != nil {
		msg := fmt.Sprintf("can't remove file %s", path)
		return e.Wrap(msg, err)
	}

	return nil
}
func (s Storage) IsExists(page *storage.Page) (bool, error) {

	fileName, err := fileName(page)
	if err != nil {
		return false, e.Wrap("can't get check if page exists", err)
	}

	path := filepath.Join(s.basePath, page.UserName, fileName)

	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		msg := fmt.Sprintf("Can't check if file %s exist", path)
		return false, e.Wrap(msg, err)
	}

	return true, nil
}
