CREATE TABLE IF NOT EXISTS "MessageHistory" ( 
				"messageID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				"chatID" INTEGER,
				"content" STRING NOT NULL,
				"sender" TEXT,
				"senderID" INTEGER NOT NULL,
				"recipient" TEXT,
				"type"		TEXT,
				"creationDate" TIMESTAMP, 
				FOREIGN KEY(chatID)REFERENCES Chats(chatID),
				FOREIGN KEY(sender)REFERENCES Users(nickName),
				FOREIGN KEY(recipient)REFERENCES Users(nickName)
				);