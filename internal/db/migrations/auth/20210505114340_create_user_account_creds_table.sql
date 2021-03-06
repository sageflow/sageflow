-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE user_account_creds (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    email citext UNIQUE NOT NULL, -- citext for case-insensitive search.
    password_hash text NOT NULL
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE user_account_creds;
