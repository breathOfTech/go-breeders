package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"go-breeders/models"
)

var testApp application

func TestMain(m *testing.M) {
	dsn := "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s"
	db, err := initMySQLDB(dsn)

	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}

	testApp = application{
		DB:     db,
		Models: *models.New(db),
	}

	os.Exit(m.Run())
}
