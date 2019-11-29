package main

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	_ "html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// JSONData it is JSON objeck what must came from outsideS
type JSONData struct {
	name      string
	fcm_token string
	viewer_id string
	viewrname string
}

var db *sql.DB

// InitDB initiate connection to database
func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

// FCMTokenSaver handler for saving FCM tokens in database
func FCMTokenSaver(w http.ResponseWriter, r *http.Request) {

}

func main() {
	var hostPort string = "127.0.0.1:8800"
	var databaseURI string = "postgres://user:pass@localhost/bookstore" //need add normal user and pass :)

	dbURI := os.Getenv("DBURI")

	if dbURI != "" {
		databaseURI = dbURI
	}

	defer InitDB(databaseURI)

	hpGlVar := os.Getenv("SERVER")

	if hpGlVar != "" {
		hostPort = hpGlVar
	}

	log.Printf("Servers started at: %v", hostPort)

	fmt.Println(http.ListenAndServe(hostPort, nil))

}
