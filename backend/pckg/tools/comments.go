package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (data *DB) Comments(w http.ResponseWriter, r *http.Request) {
	SetUpCorsResponse(w)

	// Decodes posts data into post variable
	var comment Comments

	// Decode the JSON data from the request body into the post variable
	json.NewDecoder(r.Body).Decode(&comment)
	// w.WriteHeader(http.StatusOK)
	fmt.Println("Comment is", comment, comment.PostID)

	w.Write([]byte("ok"))
	// fetches current session value
	x, err := r.Cookie("session_token")
	if err != nil {
		log.Fatal()
	}
	sessionvalue := x.Value
	// Convert data into variables for easier use
	commentId := comment.CommentID
	time := time.Now()
	content := comment.Content
	sess := data.GetSession(sessionvalue)
	postID := comment.PostID
	// Inserts post into the 'posts' table of the database
	data.CreateComment(Comments{
		// username from current session
		Author:      sess.username,
		CommentID:   commentId,
		PostID:      postID,
		Content:     content,
		CommentTime: time,
	})
}

func (data *DB) CreateComment(comment Comments) {

	fmt.Println(comment.PostID)

	stmt, err := data.DB.Prepare("INSERT INTO comments(postID, author, content, commentTime) VALUES(?, ?, ?, ?);")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(comment.PostID, comment.Author, comment.Content, comment.CommentTime)

	if err != nil {
		log.Fatal(err)
	}
}

func (data *DB) GetComments(postID int) []Comments {

	// Used to store all of the comments
	var comments []Comments

	// Used to store individual comment data
	var comment Comments

	fmt.Println("POST", postID)

	rows, err := data.DB.Query(`SELECT * FROM Comments WHERE postID =?`, postID)

	if err != nil {
		log.Fatal(err)
	}

	// Scans through every row where the postID matches the postID passed in
	for rows.Next() {

		// Populates post var with data from each post found in table
		err := rows.Scan(&comment.CommentID, &comment.PostID, &comment.Author, &comment.Content, &comment.CommentTime)

		if err != nil {
			
			fmt.Println("ERROR:", err)
			// log.Fatal(err
			return nil
		}

		// Adds each comment found from specific post to posts slice
		fmt.Println("Comment from GetComments", comment)
		comments = append([]Comments{comment}, comments...)

	}

	fmt.Println("GET commments,", comments)
	return comments

}

func (data *DB) SendComments(w http.ResponseWriter, r *http.Request) {

	SetUpCorsResponse(w)

	if r.Method == "POST" {
		var comment Comments

		json.NewDecoder(r.Body).Decode(&comment)

		fmt.Println("AFTERJSON", comment)

		comments := data.GetComments(comment.PostID)

		fmt.Println("Comment being sent", comments)

		js, err := json.Marshal(comments)

		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(js))
	}

}
