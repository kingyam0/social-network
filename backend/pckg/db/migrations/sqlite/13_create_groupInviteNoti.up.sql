CREATE TABLE GroupInviteNoti( 
	notiInviteID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	sender TEXT,
	recipient TEXT,
	FOREIGN KEY(sender)REFERENCES Users(nickName),
	FOREIGN KEY(recipient)REFERENCES GroupChat(groupID)
);