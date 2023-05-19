CREATE TABLE Users(
				userID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
				firstName TEXT NOT NULL,
				lastName TEXT NOT NULL,				
				nickName TEXT,
				email TEXT NOT NULL UNIQUE, 
				dob TEXT NOT NULL,
				aboutMe TEXT,
				avatar VARCHAR(2083),
				loggedIn TEXT,
				passwordhash TEXT NOT NULL
				);