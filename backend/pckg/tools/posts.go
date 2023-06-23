package tools

import (
	"log"
	"time"
	"net/http"
	"encoding/json"
	"fmt"
)

// Handles receiving the post data and adding it to the 'posts' table in the database
func (data *DB) Post(w http.ResponseWriter, r *http.Request) {
	SetUpCorsResponse(w)

	// Decodes posts data into post variable
	var post Post
	// Decode the JSON data from the request body into the post variable
	json.NewDecoder(r.Body).Decode(&post)
	// w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
	// fetches current session value
	x, err := r.Cookie("session_token")
	if err != nil {
		log.Fatal()
	}
	sessionvalue := x.Value
	// Convert data into variables for easier use
	postId := post.PostID
	time := time.Now()
	content := post.Content
	category := post.Category
	title := post.Title
	sess := data.GetSession(sessionvalue)
	photoUp := post.PhotoUp
	privacy := post.Privacy
	// Inserts post into the 'posts' table of the database
	data.CreatePost(Post{
		// username from current session
		Author:  sess.username,
		PostID: postId,
		Content:   content,
		Category: category,
		Title:   title,
		PhotoUp: photoUp,
		Privacy: privacy,
		CreationDate: time,
	})
}


// Handles creation of new posts
func (data *DB) CreatePost(post Post) {
	fmt.Println(&post.PhotoUp)
	stmt, err := data.DB.Prepare("INSERT INTO posts (author, title, content, category, creationDate, photoUp, privacy) VALUES (?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		log.Fatal(err)
	}
	// Uses data from post variable to insert into posts table
	_, err = stmt.Exec(&post.Author, &post.Title, &post.Content, &post.Category, &post.CreationDate, &post.PhotoUp, &post.Privacy)
	if err != nil {
		log.Fatal(err)
	}
}
// Pulls all posts from specific user and returns it as a slice of Post structs
func (data *DB) GetPosts(username string) []Post {
	// Used to store all of the posts
	var posts []Post
	// Used to store invidiual post data
	var post Post
	rows, err := data.DB.Query(`SELECT * FROM Posts WHERE author =?`, username)
	if err != nil {
		log.Fatal(err)
	}
	// Scans through every row where the username matches the username passed in
	for rows.Next() {
		// Populates post var with data from each post found in table
		err := rows.Scan(&post.PostID, &post.Author, &post.Title, &post.Content, &post.Category, &post.CreationDate, &post.PhotoUp, &post.Privacy)
		if err != nil {
			log.Fatal(err)
		}
		// Adds each post found from specific user to posts slice
		posts = append(posts, post)
	}
	return posts
}
func (data *DB) getLatestPosts() []Post {
	// Used to store all of the posts
	var posts []Post
	// Used to store invidiual post data
	var post Post
	rows, err := data.DB.Query(`SELECT * FROM Posts`)
	if err != nil {
		log.Fatal(err)
	}
	// Scans through every post
	for rows.Next() {
		// Populates post var with data from each post found in table
		err := rows.Scan(&post.PostID, &post.Author, &post.Title, &post.Content, &post.Category, &post.CreationDate, &post.PhotoUp, &post.Privacy)
		if err != nil {
			log.Fatal(err)
		}
		// Adds each post found from specific user to posts slice
		posts = append(posts, post)
	}
	return posts
}

func (data *DB) SendLatestPosts(w http.ResponseWriter, r *http.Request) {
	SetUpCorsResponse(w)
	// Send user information back to client using JSON format
	posts := data.getLatestPosts()
	// fmt.Println(userInfo)
	js, err := json.Marshal(posts)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK) // Checked in authentication.js, alerts user
	w.Write([]byte(js))
}