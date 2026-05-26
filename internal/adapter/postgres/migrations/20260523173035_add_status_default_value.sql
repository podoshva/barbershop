-- +goose Up
ALTER TABLE orders
ALTER COLUMN status SET DEFAULT 'booked';

-- +goose Down
ALTER TABLE orders
ALTER COLUMN status DROP DEFAULT;
