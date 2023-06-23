package tools

import (
	"log"
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
	var privacy string
	// Scans through each column in the 'users' row and stores the data in the variables above
	for rows.Next() {
		err := rows.Scan(&userID, &firstname, &lastname, &nickname, &email, &dob, &aboutMe, &avatar, &password, &loggedin, &privacy)
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
				Privacy:   privacy,
			},
			CreatedPosts: data.GetPosts(username),
			// You can pass POST, COMMENT, and HASHTAG etc funcs to the UserProfile struct(check old forum)
		}
	}
	return user
}

func (data *DB) GetUserProfileByID(id string) UserProfile {
	// Used to store the user's profile information
	user := UserProfile{}
	// Get a specific user's information from the 'users' table
	rows, err := data.DB.Query(`SELECT * FROM Users where userID= ?`, id)
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
	var privacy string
	// Scans through each column in the 'users' row and stores the data in the variables above
	for rows.Next() {
		err := rows.Scan(&userID, &firstname, &lastname, &nickname, &email, &dob, &aboutMe, &avatar, &password, &loggedin, &privacy)
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
				Privacy:   privacy,
			},
			// CreatedPosts: data.GetPosts(nickname),
			// You can pass POST, COMMENT, and HASHTAG etc funcs to the UserProfile struct(check old forum)
		}
	}
	return user
}

func (data *DB) GetUserID(sessionvalue string) (int, error) {
	var userID int
	err := data.DB.QueryRow("SELECT userID FROM sessions WHERE cookieValue = ?", sessionvalue).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
