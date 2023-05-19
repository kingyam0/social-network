CREATE TABLE GroupMessageHistory( 
	groupID INTEGER NOT NULL, 
	messageID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	chatID INTEGER,
	senderID INTEGER,
	sender TEXT,
	recipient TEXT,
	content TEXT,
	creationDate TIMESTAMP, 
	FOREIGN KEY(groupID)REFERENCES GroupChat(groupID),
	FOREIGN KEY(sender)REFERENCES Users(nickName),
	FOREIGN KEY(recipient)REFERENCES GroupChat(groupName)
	);