package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/antoinejaussoin/retro-board/retro-board-go/models"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// type Session struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

// var sessions []Session = []Session{
// 	Session{ID: "one", Name: "Foo "},
// 	Session{ID: "two", Name: "Bar "},
// }

func Uuid() string {
	result, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("Could not generate UUID")
	}
	return result.String()
}

func main() {
	fmt.Println("Retro-Board Backend in Go!!")

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=retroboard_go password=postgres sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println("Error while opening database", err)
		return
	}

	db.DropTableIfExists("sessions")
	db.DropTableIfExists("users")

	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Session{})

	aj := models.User{
		// ID:          Uuid(),
		Name:        "Antoine",
		AccountType: "anonymous",
		Language:    "fr",
		Username:    "antoinejaussoin",
	}

	db.Save(&aj)

	session := models.Session{
		// Id:        Uuid(),
		Url:       Uuid(),
		Name:      "my session",
		CreatedBy: aj,
	}

	db.Save(&session)

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		if err != nil {
			fmt.Println("Error on upgrade", err)
			return
		}

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		type Result struct {
			Id   string
			Name string
		}

		var results []models.Session
		db.Find(&results)
		json.NewEncoder(w).Encode(results)
		// fmt.Fprintf(w, "pong")
	})

	http.ListenAndServe("localhost:8080", nil)
}
