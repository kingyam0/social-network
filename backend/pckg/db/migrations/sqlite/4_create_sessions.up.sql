CREATE TABLE Sessions
					(
					userID INTEGER NOT NULL,
					cookieValue TEXT NOT NULL,
					firstName TEXT REFERENCES Users(firstName),
					FOREIGN KEY(userID) REFERENCES Users(userID)
					);