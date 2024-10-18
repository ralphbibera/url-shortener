package entities

import (
	"database/sql"
	"math/rand"
	"url-shortener/pkg/logger"

	"github.com/aws/aws-sdk-go/aws"
)

type Url struct {
	UrlId       *string `json:"urlId,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	Database    *sql.DB `json:"-"`
}

func (url *Url) GenerateUrlId() {
	characters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	runeSlice := make([]rune, 7)
	for i := range runeSlice {
		runeSlice[i] = characters[rand.Intn(len(characters))]
	}
	url.UrlId = aws.String(string(runeSlice))
}

func (url *Url) GetUrl() error {
	err := url.Database.QueryRow(
		`SELECT UrlId,RedirectUrl FROM url
		WHERE UrlId = $1 OR RedirectUrl = $2`,
		url.UrlId,
		url.RedirectUrl,
	).Scan(&url.UrlId, &url.RedirectUrl)
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Log("Successfully fetched existing url")
	return nil
}

func (url Url) DoesUrlExist() (bool, error) {
	var count int
	err := url.Database.QueryRow(
		`SELECT COUNT(*) FROM url 
		WHERE UrlId = $1 OR RedirectUrl = $2`,
		url.UrlId,
		url.RedirectUrl,
	).Scan(&count)
	if err != nil {
		logger.Error(err)
		return false, err
	}
	return count > 0, nil
}

func (url Url) Insert() error {
	_, err := url.Database.Exec(
		`INSERT INTO url (UrlId, RedirectUrl) VALUES ($1, $2) RETURNING *`,
		*url.UrlId,
		*url.RedirectUrl,
	)
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Log("Successfully created new url")
	return nil
}
