CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email varchar UNIQUE NOT NULL,
  fullname varchar UNIQUE NOT NULL,
  password varchar NOT NULL
);
