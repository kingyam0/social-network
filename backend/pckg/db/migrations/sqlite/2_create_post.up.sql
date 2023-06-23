CREATE TABLE Posts ( 
				"postID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
				"author" TEXT NOT NULL,
				"title" TEXT NOT NULL, 
				"content" TEXT NOT NULL, 
				"category" TEXT NOT NULL,
				"creationDate" TIMESTAMP,
				"photoUp" 	TEXT,
				"privacy"	TEXT
				);