CREATE TABLE Messages ( 
				"messageID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				"sender" TEXT REFERENCES Users(nickName),
				"recipient" TEXT REFERENCES Users(nickName),
				"content" STRING NOT NULL,
				"creationDate" TIMESTAMP
				);