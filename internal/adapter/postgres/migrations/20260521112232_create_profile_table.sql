-- +goose Up
SELECT 'up SQL query';

CREATE TABLE "profiles" (
  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "branch_id" BIGINT NOT NULL REFERENCES branches(id) ON DELETE CASCADE,
  "full_name" TEXT NOT NULL,
  "password" TEXT NOT NULL,
  "login" TEXT UNIQUE NOT NULL,
  "role" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
SELECT 'down SQL query';
