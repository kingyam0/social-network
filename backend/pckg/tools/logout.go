package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// logout handle
func (data *DB) LogoutUser(w http.ResponseWriter, r *http.Request) {
	
	if r.URL.Path != "/logout" {
		http.Error(w, "404 LOGIN PAGE NOT FOUND...", 404)
		return
	}
	
	SetUpCorsResponse(w)
	
	c, err := r.Cookie("session_token")
	if err != nil {
		log.Fatal(err)
	}
	
	sess := data.GetSession(c.Value)
	fmt.Printf("User %d wants to logout\n", sess.userID)
	
	loggedin := "false"
	
	data.DeleteSession(w, sess.userID)
	
	data.UpdateStatus(loggedin, sess.username)
	
	// Send user information back to client using JSON format
	userInfo := data.GetUserProfile(sess.username)
	
	// fmt.Println(userInfo)
	js, err := json.Marshal(userInfo)
	
	if err != nil {
		log.Fatal(err)
	}
	
	w.WriteHeader(http.StatusOK) // Checked in authentication.js, alerts user
	w.Write([]byte(js))
}
