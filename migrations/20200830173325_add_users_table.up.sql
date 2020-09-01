CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY,
  name VARCHAR(128),
  email VARCHAR(64) UNIQUE,
  hashed_password VARCHAR(64),
  role VARCHAR(16),
  created_at TIMESTAMP NOT NULL
);
