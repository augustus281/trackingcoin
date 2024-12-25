CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR UNIQUE NOT NULL,
  fullname VARCHAR UNIQUE NOT NULL,
  hashed_password TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE assets (
  id SERIAL PRIMARY KEY,
  cmc_id BIGINT NOT NULL,
  name VARCHAR UNIQUE NOT NULL,
  slug VARCHAR UNIQUE NOT NULL,
  price DOUBLE PRECISION NOT NULL, 
  percent_change_1h REAL NOT NULL,
  percent_change_24h REAL NOT NULL,
  percent_change_7d REAL NOT NULL,
  market_cap DOUBLE PRECISION NOT NULL, 
  volume_24h DOUBLE PRECISION NOT NULL, 
  circulating_supply REAL NOT NULL,
  all_time_high REAL NOT NULL,
  all_time_low REAL NOT NULL,
  turnover REAL NOT NULL,
  total_supply REAL NOT NULL,
  max_supply REAL NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX ON "assets" ("slug");