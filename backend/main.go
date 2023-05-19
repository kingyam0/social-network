package main

import (
	"fmt"
	"log"
	"net/http"

	tools "backend/pckg/tools"

	database "backend/pckg/db/migrations"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	data := database.ConnectDB("pckg/db/database.db")
	database.RunMigration("file://pckg/db/migrations/sqlite", "sqlite3://./pckg/db/database.db")
	database := &tools.DB{DB: data}

	// this handler is for the homePage
	http.HandleFunc("/", tools.HomePage)
	http.HandleFunc("/register", database.Register)
	http.HandleFunc("/login", database.Login)
	http.HandleFunc("/logout", database.LogoutUser)
	// http.HandleFunc("/posts", tools.Posts)

	// Create the hub that will manage the connections and communication with clients
	hub := tools.NewHub(database)
	go hub.Run()
	go hub.LogConns()
	// When a request is received at the "/ws" endpoint
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		database.ServeWs(hub, w, r)
	})

	fmt.Println("Starting server at port 9090")
	// here is where we listen and serve our port.x
	err2 := http.ListenAndServe(":9090", nil)
	if err2 != nil {
		log.Fatal(err2)
	}
}
