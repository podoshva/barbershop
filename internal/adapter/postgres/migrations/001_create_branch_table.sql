-- +goose Up
SELECT 'up SQL query';

CREATE TABLE "branches" (
  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "name" TEXT NOT NULL UNIQUE,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
SELECT 'down SQL query';
