CREATE TABLE IF NOT EXISTS "Chats" ( 
				"chatID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				"user1" TEXT,
				"user2" TEXT,
				"creationDate" TIMESTAMP, 
				FOREIGN KEY(user1)REFERENCES Users(nickName),
				FOREIGN KEY(user2)REFERENCES Users(nickName)
				);