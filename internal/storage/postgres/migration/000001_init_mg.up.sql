CREATE TABLE IF NOT EXISTS post (
  id SERIAL PRIMARY KEY,
  title VARCHAR(250) NOT NULL,
  text TEXT NOT NULL,
  image TEXT NOT NULL
);