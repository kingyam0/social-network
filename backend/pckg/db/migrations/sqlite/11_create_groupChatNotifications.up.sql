CREATE TABLE GroupChatNotifications( 
	groupChatNotiID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	sender TEXT,
	recipient TEXT,
	count INTEGER,
	FOREIGN KEY(sender)REFERENCES MessageHistory(sender),
	FOREIGN KEY(recipient)REFERENCES GroupChat(groupID)
	);