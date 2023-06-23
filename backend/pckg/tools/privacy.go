package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Function to update visibility in the database
func (data *DB) updatePrivacy(userID int, privacy string) {
	stmt, err := data.DB.Prepare("UPDATE Users SET privacy=? WHERE userID=?")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(privacy, userID)
}

func (data *DB) UpdateProfilePrivacy(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/updateprivacy" {
		http.Error(w, "404 LOGIN PAGE NOT FOUND...", 404)
		return
	}
	SetUpCorsResponse(w)

	if r.Method == "POST" {
		// if r.Method == "OPTIONS" {
		// 	return
		// }
		x, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "Session not found", http.StatusUnauthorized)
			return
		}
		sessionValue := x.Value
		sess := data.GetSession(sessionValue)
		var user User
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Println("error decoding json, ", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		data.updatePrivacy(sess.userID, user.Privacy)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
