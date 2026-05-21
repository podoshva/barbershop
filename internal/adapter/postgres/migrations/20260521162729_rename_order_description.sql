-- +goose Up
SELECT 'up SQL query';

ALTER TABLE orders RENAME COLUMN "desctiption" TO "description";
-- +goose Down
SELECT 'down SQL query';
