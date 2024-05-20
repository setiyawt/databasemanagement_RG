package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

type DBCredential struct {
	HostName     string
	DatabaseName string
	Username     string
	Password     string
	Port         string
}

// TODO: masukkan CAMP_ID kalian dan Credential Database kalian disini
var (
	CAMP_ID = "BE1234567" // TODO: replace this

	credential = DBCredential{
		// TODO: answer here
		HostName:     "localhost",
		DatabaseName: "BE7568661",
		Username:     "postgres",
		Password:     "postgres",
		Port:         "5432",
	}
)

func Connection() (db *sql.DB, err error) {
	//setup connection to database postgres
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		credential.HostName, credential.Port, credential.Username, credential.Password, credential.DatabaseName)

	// TODO: answer here
	db, err = sql.Open("postgres", psqlInfo)
	return

	err = db.Ping()
	if err != nil {
		return
	} else {
		fmt.Println("Successfully connected!")
		err = ioutil.WriteFile("output.txt", []byte(CAMP_ID+" "+fmt.Sprintf("%x", sha256.Sum256([]byte(CAMP_ID)))), 0644)
		return db, err
	}
}

func main() {
	_, err := Connection()

	if err != nil {
		log.Fatal(err)
	}
}
