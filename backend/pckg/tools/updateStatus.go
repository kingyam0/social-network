package tools

import "log"

func (data *DB) UpdateStatus(loggedin string, username string) {
	
	stmt, err := data.DB.Prepare("UPDATE Users SET loggedIn = ? WHERE firstName = ?;")
	
	if err != nil {
		log.Fatal(err)
	}
	
	stmt.Exec(loggedin, username)
}
