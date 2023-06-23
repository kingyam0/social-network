package tools

import (
	"net/http"
	"encoding/json"
	"log"
)

func (data *DB) CheckCookie(w http.ResponseWriter, r *http.Request) {
	SetUpCorsResponse(w)

	if r.Method == "POST" {

    var cookieValue CookieValue

    // Decode the JSON data from the request body into the comment variable
    json.NewDecoder(r.Body).Decode(&cookieValue)

    u := data.GetSession(cookieValue.CookieValue)
    userName := (u.username)

    userInfo := data.GetUserProfile(userName)

    js, err := json.Marshal(userInfo)
    if err != nil {
        log.Fatal(err)
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(js))
}
}