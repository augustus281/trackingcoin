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

CREATE TABLE notification_objects (
  id SERIAL PRIMARY KEY,
  entity_type_id BIGINT NOT NULL,
  entity_id BIGINT NOT NULL,
  status SMALLINT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE notifications (
  id SERIAL PRIMARY KEY,
  notification_object_id BIGINT NOT NULL,
  notifier_id BIGINT NOT NULL,
  status SMALLINT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT NOW(),
  CONSTRAINT fk_notification_object
    FOREIGN KEY (notification_object_id)
    REFERENCES notification_objects (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT fk_notification_notifier_id
    FOREIGN KEY (notifier_id)
    REFERENCES "users" (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);

CREATE TABLE notification_changes (
  id SERIAL PRIMARY KEY,
  notification_object_id BIGINT NOT NULL,
  actor_id BIGINT NOT NULL,
  status SMALLINT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT NOW(),
  CONSTRAINT fk_notification_object_2
    FOREIGN KEY (notification_object_id)
    REFERENCES notification_objects (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT fk_notification_actor_id
    FOREIGN KEY (actor_id)
    REFERENCES "users" (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);

CREATE TABLE user_followed_assets (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  asset_id INT REFERENCES assets(id) ON DELETE CASCADE,
  followed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON "assets" ("slug");
CREATE INDEX ON "assets" ("cmc_id");

-- Add indexes for Notification
CREATE INDEX IF NOT EXISTS fk_notification_object_idx
ON notifications (notification_object_id);

CREATE INDEX IF NOT EXISTS fk_notification_notifier_id_idx
ON notifications (notifier_id);


-- Add indexes for Notification Change
CREATE INDEX IF NOT EXISTS fk_notification_object_idx_2
ON notification_changes (notification_object_id);

CREATE INDEX IF NOT EXISTS fk_notification_actor_id_idx
ON notification_changes (actor_id);