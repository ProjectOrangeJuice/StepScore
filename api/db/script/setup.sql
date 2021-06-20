CREATE TABLE IF NOT EXISTS "userBoards" (
	"userID"	INTEGER,
	"boardID"	TEXT,
    UNIQUE("boardID","userID")
	FOREIGN KEY("userID") REFERENCES "users"("userID")
	FOREIGN KEY("boardID") REFERENCES "boards"("boardID")
);
CREATE TABLE IF NOT EXISTS "boards" (
	"boardID"	TEXT UNIQUE,
	"description"	TEXT
);
CREATE TABLE IF NOT EXISTS "steps" (
	"userID"	INTEGER,
	"date"	TEXT,
	"steps"	INTEGER,
	FOREIGN KEY("userID") REFERENCES "users"("userID")
    UNIQUE("date","userID")
);
CREATE TABLE IF NOT EXISTS "users" (
	"userID"	INTEGER NOT NULL UNIQUE,
	"username"	TEXT NOT NULL
);

