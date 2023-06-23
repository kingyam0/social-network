package tools
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	uuid "github.com/satori/go.uuid"
)
// Handles the login of existing users - validates the data and checks if it exists in the 'users' table in database
func (data *DB) Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, "404 LOGIN PAGE NOT FOUND...", 404)
		return
	}
	SetUpCorsResponse(w)
	if r.Method == "POST" {
		var user LoginData
		fmt.Println("r.Body: ", r.Body)
		err := json.NewDecoder(r.Body).Decode(&user)
		fmt.Println("USER DETAILS:", user)
		if err != nil {
			w.Write([]byte("marshalling error"))
			return
		}
		check, err := data.CheckEmailAndPass(user, w)
		if err != nil {
			fmt.Println("err2 selecting passwordhash in db by firstName or email:", err)
			if err.Error() == "sql: no rows in result set" {
				fmt.Println("on track1")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("ERROR: This username/email doesn't exist, please register to enter this forum"))
				return
			} 
				
		}
		var usID int
		if check {
			fmt.Println("SUCCESS: User logged in.")
			row := data.DB.QueryRow("SELECT userID FROM Users WHERE firstName = ?;", user.UserName)
			err := row.Scan(&usID)
			if err != nil {
				log.Fatal(err)
			}
			// Creates a new session for the user
			sess := &UserSession{}
			sess.username = user.UserName
			sess.userID = usID
			// sess.max_age = 18000
			sess.session = (uuid.NewV4().String() + "&" + strconv.Itoa(sess.userID) + "&" + sess.username)
			user.LoggedIn = "true"
			// Set client cookie for "session_token" as session token we just generated, also set expiry time to 120 minutes
			http.SetCookie(w, &http.Cookie{
				Name:     "session_token",
				Value:    sess.session,
				// Expires:  time.Now().Add(24 * time.Hour),
				Secure:   true,
				SameSite: http.SameSiteNoneMode,
			})
			// Insert data into session variable
			data.InsertSession(*sess)
			data.UpdateStatus(user.LoggedIn, user.UserName)
			// Send user information back to client using JSON format
			userInfo := data.GetUserProfile(user.UserName)
			js, err := json.Marshal(userInfo)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK) // Checked in authentication.js, alerts user
			w.Write([]byte(js))
		}
	}
	if r.Method == "OPTIONS" {
		return
	}
}
