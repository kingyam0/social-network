CREATE TABLE IF NOT EXISTS "Comments" ( 
				"commentID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
				"postID" INTEGER NOT NULL,
				"authorID" INTEGER NOT NULL,
				"author" TEXT NOT NULL,
				"content" TEXT NOT NULL, 
				"creationDate" TIMESTAMP,
				FOREIGN KEY(postID)REFERENCES posts(postID),
				FOREIGN KEY(authorID)REFERENCES users(userID)
				);