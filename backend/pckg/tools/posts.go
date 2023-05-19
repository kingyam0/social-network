package tools

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"
// )

// var thePost Post

// var data *DB

// func Posts(w http.ResponseWriter, r *http.Request) {

// 	var postTime = time.Now()

// 	err := json.NewDecoder(r.Body).Decode(&thePost)

// 	fmt.Println("=================", err)
// 	fmt.Println(thePost, "WERE ARE GETTING HERE-=-=-=-=-=-=-=-=")

// 	fmt.Println(CookieID, "this is the cookie I@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@'")
// 	var usr = GetUserByCookie(CookieID)

// 	var usrID = usr.UserID

// 	_, err = data.DB.Exec(`INSERT INTO Posts (
// 		authorID,
// 		author,
// 		title,
// 		content,
// 		category,
// 		creationDate,
// 		cookieID
// 		) VALUES(?,?,?,?,?,?,?)`, usrID, usr.NickName, thePost.Title, thePost.Content, thePost.Category, postTime, thePost.Cookie)
// 	if err != nil {
// 		fmt.Println("Error inserting into 'Posts' table: ", err)
// 		return
// 	}
// }

// func GetPosts() []Post {
// 	var posts []Post
// 	var myPost Post

// 	rows, errPost := data.DB.Query("SELECT postID, author, category, title, content, creationDate FROM Posts;")
// 	if errPost != nil {
// 		fmt.Println("Error retrieving posts from database: \n", errPost)
// 		return nil
// 	}

// 	for rows.Next() {
// 		//copy row columns into corresponding variables
// 		err := rows.Scan(&myPost.PostID, &myPost.Author, &myPost.Category, &myPost.Title, &myPost.Content, &myPost.PostTime)
// 		if err != nil {
// 			fmt.Println("error copying post data: ", err)
// 		}

// 		//aggregate all posts separated by '\n'
// 		posts = append(posts, myPost)
// 	}
// 	rows.Close()

// 	return posts
// }

// // To send all posts to front-end via http handle: "/getPosts"
// func SendLatestPosts(w http.ResponseWriter, r *http.Request) {
// 	//Send user information back to client using JSON format
// 	posts := GetPosts()
// 	js, err := json.Marshal(posts)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	w.WriteHeader(http.StatusOK) //Ceck in authentication.js, alerts user
// 	w.Write([]byte(js))
// }

// func (p *Post) Modify(ck string) {
// 	p.Cookie = ck
// }

// // GetUserByCookie ...
// func GetUserByCookie(cookieValue string) *User {
// 	var userID int64

// 	if err := data.DB.QueryRow("SELECT userID from Sessions WHERE cookieValue = ?", cookieValue).Scan(&userID); err != nil {
// 		fmt.Println("++++++++++++++++++++++++___________________________===", cookieValue)
// 		return nil
// 	}
// 	u := FindByUserID(userID)
// 	return u
// }

// // function for new user
// func NewUser() *User {
// 	return &User{}
// }

// // Find the user by their ID
// func FindByUserID(UID int64) *User {
// 	u := NewUser()
// 	if err := data.DB.QueryRow("SELECT userID, firstName, lastName, nickName, age, gender, email, passwordhash FROM Users WHERE userID = ?", UID).
// 		Scan(&u.UserID, &u.FirstName, &u.LastName, &u.NickName, &u.AboutMe, &u.Avatar, &u.Email); err != nil {
// 		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++error FindByUserID: ", err)
// 		return nil
// 	}
// 	return u
// }
