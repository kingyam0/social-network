package tools

import (
	"database/sql"
	"time"
)

type DB struct {
	DB *sql.DB
}
type RegisterData struct {
	UserID    int
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	NickName  string `json:"NickName"`
	DOB       string `json:"DOB"`
	Email     string `json:"Email"`
	Avatar	  string `json:"Avatar"`
	AboutMe   string `json:"AboutMe"`
	Access    int    // 0 means no access, not logged in
	LoggedIn  string
	Posts     []Post
	Comments  []Comments
	Password  string `json:"PassWord"`
}
type User struct {
	UserID    int
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	NickName  string `json:"NickName"`
	DOB       string `json:"DOB"`
	Avatar	  string `json:"Avatar"`
	AboutMe   string `json:"AboutMe"`
	Email     string `json:"Email"`
	LoggedIn  string
}
type UserProfile struct {
	User User
	// CreatedPosts  []Post
	// Hashtags      []Hashtag
	// Notifications []Notifications
	// Messages []Message
}
type GroupChat struct {
	GroupID     int
	GroupName   string    `json:"GroupName"`
	GroupDes    string    `json:"GroupDescription"`
	Creator     string    // creator
	CreatedTime time.Time `json:"CreatedTime"`
}
type LoginData struct {
	UserName string `json:"Username"`
	Password string `json:"Password"`
	LoggedIn string
}

// each session contains the username of the user and the time at which it expires
type UserSession struct {
	username string
	userID   int
	session  string
	max_age  int
}
type Cookie struct {
	Name    string
	Value   string
	Expires time.Time
}
type Comments struct {
	CommentID      int
	Author         string    `json:"Author"`
	PostID         int       `json:"PstID"`
	Content        string    `json:"CommContent"`
	CommentTime    time.Time `json:"CommentTime"`
	CommentTimeStr string
}
type Post struct {
	PostID      int       `json:"PostID"`
	Author      string    // author
	Title       string    `json:"PostTitl"`
	Content     string    `json:"PostCont"`
	Category    string    `json:"PostCat"`
	PostTime    time.Time `json:"PostTime"`
	PostTimeStr string
	Comments    []Comments
	IPs         string
	Public 		string
}
type GroupPost struct {
	GroupID     int       `json:"GroupID"`
	GroupPostID int       `json:"GPostID"`
	Author      string    // author
	Title       string    `json:"GPostTitl"`
	Content     string    `json:"GPostCont"`
	Category    string    `json:"GPostCat"`
	PostTime    time.Time `json:"GPostTime"`
	PostTimeStr string
	Comments    []Comments
	IPs         string
	Public 		string

}
type Message struct {
	MessageID int
	ChatID    int
	SenderID  string `json:"SenderID"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"Messagecontent"`
	Type      string
	Date      time.Time `json:"date"`
}

type GroupMessage struct {
	GroupID   int 
	MessageID int
	ChatID    int
	SenderID  string `json:"SenderID"`
	Sender    string `json:"sender"`
	Recipient string `json:"Grouprecipient"`
	Content   string `json:"GroupChatcontent"`
	Type      string
	Date      time.Time `json:"date"`
}

type Chat struct {
	ChatID int
	User1  string
	User2  string
	Date   string
}
type ChatHistoryCheck struct {
	ChatID     int
	ChatExists bool
	// ChatHistory []Message
}
type Notification struct {
	NotificationID        int
	NotificationSender    string `json:"notificationsender"`
	NotificationRecipient string `json:"notificationrecipient"`
	NotificationCount     int    `json:"notificationcount"`
	NotificationSeen      string `json:"notificationseen"`
}
type NotificationCheck struct {
	NotificationID int
	NotifExists    bool
}
