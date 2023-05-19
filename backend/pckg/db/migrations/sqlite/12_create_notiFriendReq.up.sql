CREATE TABLE NotiFriendRequest( 
	notiID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	sender TEXT,
	recipient TEXT,
	FOREIGN KEY(sender)REFERENCES Users(nickName),
	FOREIGN KEY(recipient)REFERENCES Users(userID)
);