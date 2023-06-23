package tools

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

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
