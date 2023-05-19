package tools

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// instance of the database
// var DB *sql.DB

// instance of the User struct

func SetUpCorsResponse(w http.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
}
// function that defines what happens when user registers to our social network
func (data *DB) Register(w http.ResponseWriter, r *http.Request) {
	// if Path is not /register then return error
	if r.URL.Path != "/register" {
		http.Error(w, "404 REGISTRATION PAGE NOT FOUND...", 404)
		return
	}

	SetUpCorsResponse(w)

	var LiveUser *RegisterData

	// read the body of the page and return bytes
	fmt.Println("r.Body: ", r.Body)
	
	err := json.NewDecoder(r.Body).Decode(&LiveUser)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("marshalling error"))
		return
	}

	// unmarshall bytes received by json & the LiveUser
	// json.Unmarshal(bytes, &LiveUser)
	fmt.Println("LiveUser: ", LiveUser)
	// variable hash to store password as hash
	var hash []byte
	password := LiveUser.Password
	hash, err2 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err2 != nil {
		// Status Not Acceptable=406
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Println("BCRYPT ERROR 2 LINE 47 REGISTER.GO")
		return
	}
	_, err = data.DB.Exec(`INSERT INTO Users (
		firstName,
		lastName,
		nickName,
		email,
		dob,
		aboutMe,
		avatar,
		passwordhash
	)VALUES(?,?,?,?,?,?,?,?)`, LiveUser.FirstName, LiveUser.LastName, LiveUser.NickName, LiveUser.Email, LiveUser.DOB, LiveUser.AboutMe, LiveUser.Avatar, hash)
	
	if err != nil {
		// convey Status Bad Request
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("ERROR WITH INSERTING INTO TABLE:", err)
		if err.Error() == "UNIQUE constraint failed: Users.email" {
			w.Write([]byte("ERROR: This email already exists, please log in instead"))
		} else if err.Error() == "UNIQUE constraint failed: Users.nickName" {
			w.Write([]byte("ERROR: This username already exists, please log in instead"))
		}
		// w.Write([]byte(err.Error()))
		return
	}
	
	js, _ := json.Marshal(&LiveUser)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(js))
	fmt.Println(js)
		
	if r.Method == "OPTIONS" {
		return
	}

	
}
