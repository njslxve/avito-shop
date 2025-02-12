CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  username TEXT UNIQUE NOT NULL,
  pass TEXT NOT NULL,
  coins BIGINT NOT NULL DEFAULT 1000
);

CREATE TABLE IF NOT EXISTS items (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  type TEXT NOT NULL,
  price BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS item_transactions (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id),
  item_id UUID NOT NULL REFERENCES items(id),
  ceated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_transactions (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  from_user_id UUID NOT NULL REFERENCES users(id),
  to_user_id UUID NOT NULL REFERENCES users(id),
  amount BIGINT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


INSERT INTO items (type, price) VALUES
  ('t-shirt', 80),
  ('cup', 20),
  ('book', 50),
  ('pen', 10),
  ('powerbank', 200),
  ('hoody', 300),
  ('umbrella', 200),
  ('socks', 10),
  ('wallet', 50),
  ('pink-hoody', 500);