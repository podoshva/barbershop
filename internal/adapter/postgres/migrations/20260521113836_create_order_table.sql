-- +goose Up
SELECT 'up SQL query';

CREATE TABLE "orders" (
  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "profile_id" BIGINT NOT NULL REFERENCES profiles(id) ON DELETE CASCADE,
  "branch_id" BIGINT NOT NULL,
  "date" TIMESTAMPTZ NOT NULL,
  "customer_phone" TEXT NOT NULL,
  "desctiption" TEXT NOT NULL,
  "status" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose Down
SELECT 'down SQL query';
