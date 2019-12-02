package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// JSONData it is JSON objeck what must came from outsideS
type JSONData struct {
	Name       *string `json:"name"`
	FCMToken   *string `json:"fcm_token"`
	ViewerID   *string `json:"viewer_id"`
	ViewerName *string `json:"viewer_name"`
}

//JSONDataDetail detail representation of Viewer device
type JSONDataDetail struct {
	ID         int     `json:"id"`
	Name       *string `json:"name"`
	FCMToken   *string `json:"fcm_token"`
	ViewerID   *string `json:"viewer_id"`
	ViewerName *string `json:"viewer_name"`
}

var db *sql.DB

// Env global structure for save all neaded global variables
type Env struct {
	db *sql.DB
}

//DBConnection TODO work on function DB connection and close the connection use 'defer'
func DBConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "fcm_t.db")
	if err != nil {
		log.Panic(err)
		// errors.New("Db connection not exist!")
	}
	return db, nil
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
	// var databaseURI string = "postgres://user:pass@localhost/bookstore" //need add normal user and pass :)

	// dbURI := os.Getenv("DBURI")

	// if dbURI != "" {
	// 	databaseURI = dbURI
	// }

	// db, err := InitDB(databaseURI)

	db, err := DBConnection()

	if err != nil {
		log.Panicf("Database connection error! %v", err)
	}

	env := &Env{db: db}

	hpGlVar := os.Getenv("SERVER")

	if hpGlVar != "" {
		hostPort = hpGlVar
	}

	log.Printf("Servers started at: %v", hostPort)
	http.HandleFunc("/", env.CreateFCMToken)
	fmt.Println(http.ListenAndServe(hostPort, nil))

}

// CreateFCMToken handler for saving FCM tokens in database
func (env Env) CreateFCMToken(w http.ResponseWriter, r *http.Request) {
	// TODO finish   https://www.alexedwards.net/blog/organising-database-access
	decoder := json.NewDecoder(r.Body)
	var t JSONData
	err := decoder.Decode(&t)
	// var p JSONData
	data, _ := json.Marshal(t)

	if err != nil {
		log.Printf("Can not read JSON from BODY! %v", err)
		http.Error(w, "Invalid JSON!", http.StatusBadRequest)
		return
	}
	log.Printf("Accepted JSON: %s", data)

	rows, err := db.Exec("INSERT INTO devices (name, fcm_token, viewer_id, viewer_name) VALUES ($1, $2, $3, $4)", t.Name, t.FCMToken, t.ViewerID, t.ViewerName)

	if err != nil {
		log.Printf("Error is raiced! %v", err)
	}
	for rows.Next {
		log.Printf("ID: %v, FCM: %v, viewer: %v, Viewer_name %v", rows.ID, rows.Name)

	}
}

// UpdateFCMToken handler for saving FCM tokens in database
func UpdateFCMToken(w http.ResponseWriter, r *http.Request) {

}

// DeleteFCMToken handler for saving FCM tokens in database
func DeleteFCMToken(w http.ResponseWriter, r *http.Request) {

}
