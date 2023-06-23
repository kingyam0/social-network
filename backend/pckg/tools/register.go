package tools

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)


// function that defines what happens when user registers to our social network
func (data *DB) Register(w http.ResponseWriter, r *http.Request) {
	// if Path is not /register then return error
	if r.URL.Path != "/register" {
		http.Error(w, "404 REGISTRATION PAGE NOT FOUND...", 404)
		return
	}

	SetUpCorsResponse(w)

	if r.Method == "POST"{
		var LiveUser *RegisterData
		// read the body of the page and return bytes
		fmt.Println("r.Body: ", r.Body)
		
		err := json.NewDecoder(r.Body).Decode(&LiveUser)
		
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("marshalling error"))
			return
		}

		if err != nil {
			fmt.Println(err)
			w.Write([]byte("marshalling error"))
			return
		}

		fmt.Println("LiveUser: ", LiveUser)


		//handling empty data being populated into DB
		if ( LiveUser.FirstName == "" || LiveUser.LastName == "" || LiveUser.Email == "" || LiveUser.DOB == "" || LiveUser.Password == "") {
			fmt.Println("reg-user data empty")
			return
		}

		LiveUser.LoggedIn = "false"

		var hash []byte
		//variable hash stores the password as hash
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
			loggedIn,
			passwordhash
		)VALUES(?,?,?,?,?,?,?,?,?)`, LiveUser.FirstName, LiveUser.LastName, LiveUser.NickName, LiveUser.Email, LiveUser.DOB, LiveUser.AboutMe, LiveUser.Avatar, LiveUser.LoggedIn, hash)
		
		if err != nil {
			// convey Status Bad Request
			fmt.Println("ERROR WITH INSERTING INTO TABLE:", err)
			if err.Error() == "UNIQUE constraint failed: Users.email" {
				w.Write([]byte("ERROR: This email already exists, please log in instead"))
			} 
			return
		}
		
		// js, _ := json.Marshal(&LiveUser)
		// w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		// w.Write([]byte(js))
		// fmt.Println(js)
			
		// if r.Method == "OPTIONS" {
			// return
		// }
	}

	
}
