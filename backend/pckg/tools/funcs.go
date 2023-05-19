package tools

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Pulls specific user's data and posts data from database and returns it as a User struct
func (data *DB) GetUserProfile(username string) UserProfile {

	// Used to store the user's profile information
	user := UserProfile{}
	// Get a specific user's information from the 'users' table

	rows, err := data.DB.Query(`SELECT * FROM Users where firstName= ?`, username)

	if err != nil {
		log.Fatal(err)
	}

	// Used to store the user's data so we can add it to struct later on
	var userID int
	var firstname string
	var lastname string
	var email string
	var nickname string
	var password string
	var dob string
	var aboutMe string
	var avatar string
	var loggedin string

	// Scans through each column in the 'users' row and stores the data in the variables above
	for rows.Next() {
		err := rows.Scan(&userID, &firstname, &lastname, &nickname, &email, &dob, &aboutMe, &avatar, &password, &loggedin)

		if err != nil {
			log.Fatal(err)
		}

		// This contains the specific user's data as well as all of their posts
		user = UserProfile{
			User: User{
				UserID:    userID,
				Avatar:    avatar,
				NickName:  nickname,
				FirstName: firstname,
				LastName:  lastname,
				Email:     email,
				AboutMe:   aboutMe,
				LoggedIn:  loggedin,
			},

			// You can pass POST, COMMENT, and HASHTAG etc funcs to the UserProfile struct(check old forum)
		}
	}
	return user
}

func (data *DB) CheckEmailAndPass(user LoginData, w http.ResponseWriter) (bool, error) {

	// Only set to true if the email/username IS found in the database
	emailPassCombinationValid := false
	userPassCombinationValid := false

	// Checks if user entered an email or username
	enteredEmail := strings.Contains(user.UserName, "@")

	// EMAIL CHECK
	if enteredEmail {

		// Checks if email/pass combination exists in database
		var passwordHash string

		row := data.DB.QueryRow("SELECT passwordhash FROM Users WHERE email = ?", user.UserName)
		err := row.Scan(&passwordHash)

		if err != nil {
			fmt.Println("Error with password hash:", err)
			return false, err
		}

		// If the password hash matches the password entered, the email/pass combination is valid
		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))

		if err == nil {
			emailPassCombinationValid = true
		}
	} else {
		// Checks if username/pass combination exists in database
		var passwordHash string

		row := data.DB.QueryRow("SELECT passwordhash FROM Users WHERE firstName = ?", user.UserName)
		err := row.Scan(&passwordHash)

		if err != nil {

			if err == sql.ErrNoRows {
				return false, fmt.Errorf("user not found with nickname %s", user.UserName)
			}

			fmt.Println("Error with password hash:", err)

			return false, err
		}

		// If the password hash matches the password entered, the user/pass combination is valid
		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))

		if err == nil {
			userPassCombinationValid = true
		}
	}

	return emailPassCombinationValid || userPassCombinationValid, nil
}

func (data *DB) UpdateStatus(loggedin string, username string) {

	stmt, err := data.DB.Prepare("UPDATE Users SET loggedIn = ? WHERE firstName = ?;")

	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec(loggedin, username)
}
