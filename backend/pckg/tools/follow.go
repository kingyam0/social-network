package tools

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

// The "target user" refers to the user who is being followed or the user who is the recipient of the follow request
func (data *DB) SaveFollowRequest(request FollowRequest) error {
	
	stmt, err := data.DB.Prepare("INSERT INTO FollowRequests (userID, target_user_id, request_date) VALUES (?, ?, ?)")
	
	if err != nil {
		return err
	}
	
	_, err = stmt.Exec(request.UserID, request.TargetUserID, time.Now())
	
	if err != nil {
		return err
	}
	
	return nil
}

func (data *DB) AcceptFollowRequest(userID, targetUserID int) error {

	// Check if a follow request exists between the users
	exists, err := data.FollowRequestExists(userID, targetUserID)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("follow request does not exist")
	}

	// Update the followers and following counts
	err = data.IncrementFollowersCount(targetUserID)

	if err != nil {
		return err
	}

	err = data.IncrementFollowingCount(userID)

	if err != nil {
		return err
	}

	// Delete the follow request
	err = data.DeleteFollowRequest(userID, targetUserID)

	if err != nil {
		return err
	}

	return nil
}

func (data *DB) IncrementFollowersCount(userID int) error {

	stmt, err := data.DB.Prepare("UPDATE followers SET followers = followers + 1 WHERE user_id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID)

	if err != nil {
		return err
	}

	return nil
}

func (data *DB) IncrementFollowingCount(userID int) error {

	stmt, err := data.DB.Prepare("UPDATE followers SET following = following + 1 WHERE user_id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID)

	if err != nil {
		return err
	}

	return nil
}

func (data *DB) DeleteFollowRequest(userID, targetUserID int) error {

	stmt, err := data.DB.Prepare("DELETE FROM follow_requests WHERE user_id = ? AND target_user_id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID, targetUserID)

	if err != nil {
		return err
	}
	return nil
}

func (data *DB) FollowRequestExists(userID, targetUserID int) (bool, error) {

	var count int

	err := data.DB.QueryRow("SELECT COUNT(*) FROM follow_requests WHERE user_id = ? AND target_user_id = ?", userID, targetUserID).Scan(&count)

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (data *DB) SendFollowData(w http.ResponseWriter, r *http.Request) {

	SetUpCorsResponse(w)

	var follow FollowRequest

	json.NewDecoder(r.Body).Decode(&follow)
	w.Write([]byte("follow ok"))

	// feches current session value
	x, err := r.Cookie("session_token")

	if err != nil {
		log.Fatal()
	}

	sessionvalue := x.Value

	userID, err := data.GetUserID(sessionvalue)

	if err != nil {
		log.Fatal(err)
	}

	follow.UserID = userID

	data.SaveFollowRequest(follow)
}
