ALTER TABLE posts ADD COLUMN author_id INTEGER REFERENCES users(id);