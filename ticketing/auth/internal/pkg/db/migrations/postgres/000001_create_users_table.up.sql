CREATE TABLE IF NOT EXISTS users
( user_id SERIAL PRIMARY KEY,
  email VARCHAR (255) UNIQUE NOT NULL,
  password VARCHAR (255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP);