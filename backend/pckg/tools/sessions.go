package tools

import (
	"fmt"
	"log"
	"net/http"
)

// Inserts session into sessions table
func (data *DB) InsertSession(sess UserSession) {
	
	stmnt, err := data.DB.Prepare("INSERT INTO Sessions (cookieValue, userID, firstName) VALUES (?, ?, ?)")
	
	if err != nil {
		fmt.Println("Error inserting session into table:", err)
	}
	fmt.Println("Are we getting here////////////////////////////////////////////////////////////////////////////////////////////")
	defer stmnt.Close()
	
	stmnt.Exec(sess.session, sess.userID, sess.username)
}

// TODO: Clarification
// User's cookie expires when browser is closed, delete the cookie from the database.
func (data *DB) DeleteSession(w http.ResponseWriter, userID int) error {
	
	cookie := &http.Cookie{
		Name:   "session_token",
		Value:  "",
		MaxAge: -1,
	}
	
	http.SetCookie(w, cookie)
	
	stmt, err := data.DB.Prepare("DELETE FROM Sessions WHERE userID = ?;")
	
	// defer stmt.Close()
	stmt.Exec(userID)
	
	if err != nil {

		log.Fatal(err)

		fmt.Println("DeleteSession err: ", err)
		
		return err
	}
	
	return nil
}

// Checks all sessions from sessions table and returns latest session
func (data *DB) GetSession(cookie string) UserSession {
	
	// Used to store session data
	session := UserSession{}
	
	// Checks all sessions from sessions table
	rows, err := data.DB.Query(`SELECT * FROM Sessions WHERE cookieValue = ?;`, cookie)
	
	if err != nil {
		log.Fatal(err)
	}
	
	// Used to store individual session data
	var userID int
	var cookieValue string
	var nickname string
	
	// For each session found, populate the variable above
	for rows.Next() {
		
		err2 := rows.Scan(&userID, &cookieValue, &nickname)
		
		if err2 != nil {
			log.Fatal(err2)
		}
		
		// Overwrites every session, leaving only data for the latest session
		session = UserSession{
			userID:   userID,
			session:  cookieValue,
			username: nickname,
		}
	}
	return session
}
