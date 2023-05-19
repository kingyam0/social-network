CREATE TABLE IF NOT EXISTS "Posts" ( 
				"postID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
				"authorID" INTEGER NOT NULL,
				"author" TEXT NOT NULL,
				"title" TEXT NOT NULL, 
				"content" TEXT NOT NULL, 
				"category" TEXT NOT NULL,
				"creationDate" TIMESTAMP,
				"post" TEXT,
				FOREIGN KEY(authorID)REFERENCES users(userID),
				FOREIGN KEY(author)REFERENCES users(firstName)
				);