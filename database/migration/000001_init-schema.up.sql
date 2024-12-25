CREATE TYPE notification_type AS ENUM ('email', 'discord', 'telegram');

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR UNIQUE NOT NULL,
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

CREATE TABLE notifications (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  noti_type notification_type NOT NULL DEFAULT 'email', 
  message TEXT,
  is_read BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_followed_assets (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  asset_id INT REFERENCES assets(id) ON DELETE CASCADE,
  followed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON "assets" ("slug");
CREATE INDEX ON "assets" ("cmc_id");