package rds

import (
	"log"
	"testing"
	"url-shortener/pkg/entities"
	"url-shortener/pkg/logger"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}
}

func TestConnect(t *testing.T) {
	database, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	err = database.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPing(t *testing.T) {
	database, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	log.Print("Ping!")
	err = database.Ping()
	if err != nil {
		t.Fatal(err)
	}
	log.Println("Pong!")
}

func TestInitializeTable(t *testing.T) {
	database, err := Connect()
	if err != nil {
		t.Error(err)
	}
	defer database.Close()

	_, err = database.Exec(`DROP TABLE IF EXISTS url`)
	if err != nil {
		t.Error(err)
	}

	_, err = database.Exec(
		`CREATE TABLE url (
			urlid VARCHAR(7) PRIMARY KEY,
			redirecturl TEXT
		)`,
	)
	if err != nil {
		t.Error(err)
	}
}

func TestQueryAll(t *testing.T) {
	database, err := Connect()
	if err != nil {
		t.Error(err)
	}
	defer database.Close()

	rows, err := database.Query(`SELECT UrlId,RedirectUrl FROM url`)
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()

	var urls []entities.Url

	for rows.Next() {
		var url entities.Url

		err := rows.Scan(&url.UrlId, &url.RedirectUrl)
		if err != nil {
			t.Error(err)
		}

		urls = append(urls, url)
	}

	if err := rows.Err(); err != nil {
		t.Error(err)
	}

	logger.Log(urls)
}

func TestGetItem(t *testing.T) {
	database, err := Connect()
	if err != nil {
		t.Error(err)
	}
	defer database.Close()

	var url entities.Url
	err = database.QueryRow(
		`SELECT UrlId,RedirectUrl FROM url
		WHERE UrlId = $1`,
		"Rb57PeW",
	).Scan(&url.UrlId, &url.RedirectUrl)
	if err != nil {
		t.Error(err)
	}

	logger.Log(url)
}

func TestInsertItem(t *testing.T) {
	database, err := Connect()
	if err != nil {
		t.Error(err)
	}
	defer database.Close()

	url := entities.Url{
		RedirectUrl: aws.String("www.google.com"),
	}
	url.GenerateUrlId()

	rows, err := database.Exec(
		`INSERT INTO url (UrlId, RedirectUrl) VALUES ($1, $2) RETURNING *`,
		*url.UrlId,
		*url.RedirectUrl,
	)
	if err != nil {
		t.Error(err)
	}
	logger.Log(rows)
}
