-- +goose Up
CREATE TABLE promos (
  id SERIAL PRIMARY KEY,
  user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  promo_code VARCHAR(10) NOT NULL,
  amount DOUBLE PRECISION NOT NULL DEFAULT 0,
  start_date DATE NOT NULL,
  end_date DATE NOT NULL
);

-- +goose Down
DROP TABLE promos;
