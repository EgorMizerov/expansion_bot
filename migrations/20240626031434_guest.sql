-- +goose Up
CREATE TABLE guests (
    phone_number text PRIMARY KEY,
    telegram_id int NOT NULL
);

-- +goose Down
DROP TABLE guests;
