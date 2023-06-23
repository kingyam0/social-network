package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)
func 

func (data *DB) SendAllMessages(username, recipient string) []Chat {
	var chat Chat
	var conversation []Chat
	rows, err := data.DB.Query(`SELECT sender, recipient, content, creationDate FROM Messages where (sender = ? AND recipient = ?) OR (sender = ? AND recipient = ?)`, username, recipient, recipient, username)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("after sql:", rows)
	for rows.Next() {
		// fmt.Println("inside query loop")
		err := rows.Scan(&chat.MessageSender, &chat.MessageRecipient, &chat.Content, &chat.CreatedAt)
		if err != nil {
			log.Fatal("conversation error", err)
		}
		// fmt.Println("Messages", username, recipient, ":", &loading.Message)
		conversation = append(conversation, chat)
	}
	// fmt.Println("Con", conversation)
	return conversation
}

func (data *DB) SaveChat(chat Chat) Chat {
	stmnt, err := data.DB.Prepare("INSERT INTO Messages (sender, recipient ,content, creationDate) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmnt.Exec(chat.MessageSender, chat.MessageRecipient, chat.Content, chat.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	return chat
}

func (data *DB) Chat(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/chat" {
		http.Error(w, "404 CHAT PAGE NOT FOUND...", 404)
		return
	}
	SetUpCorsResponse(w)

	if r.Method == "POST" {
		var chat Chat

		json.NewDecoder(r.Body).Decode(&chat)
		w.Write([]byte("chat ok"))
		// feches current session value
		x, err := r.Cookie("session_token")
		if err != nil {
			log.Fatal()
		}
		sessionvalue := x.Value
		content := chat.Content
		time := time.Now()
		recipient := chat.MessageRecipient
		sess := data.GetSession(sessionvalue)

		chat = data.SaveChat(Chat{
			MessageSender:    sess.username,
			MessageRecipient: recipient,
			Content:          content,
			CreatedAt:        time,
		})
		fmt.Println("<---------------------------------------->")
		fmt.Println("SENDER: ", chat.MessageSender)
		fmt.Println("RECEPIENT: ", chat.MessageRecipient)
		fmt.Println("MESSAGE Content: ", chat.Content)
		fmt.Println("TIME: ", chat.CreatedAt)
		fmt.Println("<---------------------------------------->")

	}
}

func (data *DB) Messages(w http.ResponseWriter, r *http.Request) {
	var loading AllMessage
	err := json.NewDecoder(r.Body).Decode(&loading)
	if err != nil {
		log.Fatal("Messages handler error: ", err)
	}
	conv := data.SendAllMessages(
		loading.SendersUsername,
		loading.RecipientsUsername,
	)
	// data.DeleteNotification(loading.RecipientsUsername, loading.SendersUsername)
	js, err := json.Marshal(conv)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(js))
}
