CREATE TABLE Comments ( 
				commentID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
				postID INTEGER NOT NULL,
				author TEXT NOT NULL,
				content TEXT NOT NULL, 
				commentTime TIMESTAMP,
				FOREIGN KEY(postID)REFERENCES posts(postID)
				);