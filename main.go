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

// var db *sql.DB

// Env global structure for save all neaded global variables
type Env struct {
	db *sql.DB
}

// InitDB initiate connection to database
func InitDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// 	var err error
// 	// env = &Env{}
// 	db, err = sql.Open("postgres", dataSourceName)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	if err = db.Ping(); err != nil {
// 		log.Panic(err)
// 	}
// }

func main() {
	var hostPort string = "127.0.0.1:8800"
	var databaseURI string = "postgres://user:pass@localhost/bookstore" //need add normal user and pass :)

	dbURI := os.Getenv("DBURI")

	if dbURI != "" {
		databaseURI = dbURI
	}

	db, err := InitDB(databaseURI)

	if err != nil {
		log.Panicf("Database connection error! %v", err)
	}

	env := &Env{db: db}

	hpGlVar := os.Getenv("SERVER")

	if hpGlVar != "" {
		hostPort = hpGlVar
	}

	log.Printf("Servers started at: %v", hostPort)
	http.Handle("/", CreateFCMToken(env))
	fmt.Println(http.ListenAndServe(hostPort, nil))

}

// CreateFCMToken handler for saving FCM tokens in database
func CreateFCMToken(w http.ResponseWriter, r *http.Request, env *Env) {
	// TODO finish   https://www.alexedwards.net/blog/organising-database-access
}

// UpdateFCMToken handler for saving FCM tokens in database
func UpdateFCMToken(w http.ResponseWriter, r *http.Request) {

}

// DeleteFCMToken handler for saving FCM tokens in database
func DeleteFCMToken(w http.ResponseWriter, r *http.Request) {

}
