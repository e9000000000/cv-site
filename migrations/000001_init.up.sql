CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    username TEXT UNIQUE,
    password_hash TEXT,
    token TEXT,
    is_admin BOOLEAN
);

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    title TEXT,
    text TEXT
);
