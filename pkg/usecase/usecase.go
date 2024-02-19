package usecase

import (
	"errors"
	"math/rand"
	"time"

	"github.com/url-shortner/pkg/config"
	interfaceRepo "github.com/url-shortner/pkg/repository/interface"
	interfaceUsecase "github.com/url-shortner/pkg/usecase/interface"
)

type Usecase struct {
	repo   interfaceRepo.Repository
	config config.Config
}

func NewUseCase(repo interfaceRepo.Repository, config config.Config) interfaceUsecase.Usecase {
	return &Usecase{repo: repo, config: config}
}

func (u *Usecase) Shorten(url string) (string, error) {
	const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6
	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charSet[rand.Intn(len(charSet))]
	}
	err := u.repo.Save(url, string(shortKey))
	if err != nil {
		return "", errors.New("error in saving url")
	}

	shortenedURL := u.config.BASE_URL + "/short/" + string(shortKey)
	return shortenedURL, nil
}

func (u *Usecase) Redirect(shorten string) (string, error) {
	exist, err := u.repo.IsShortExist(shorten)
	if err != nil {
		return "", errors.New("error in fetching data")
	}
	if !exist {
		return "", errors.New("this url not found")
	}
	url, err := u.repo.GetURL(shorten)
	if err != nil {
		return "", errors.New("error in fetching data")
	}
	err = u.repo.UpdateCount(shorten)
	if err != nil {
		return "", errors.New("error in fetching data")
	}
	return url, nil
}

func (u *Usecase) Count(shorten string) (int, error) {
	count, err := u.repo.GetCount(shorten)
	if err != nil {
		return 0, errors.New("error in fetching data")
	}
	return count, nil
}
