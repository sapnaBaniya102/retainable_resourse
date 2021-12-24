
-- +migrate Up
CREATE SEQUENCE IF NOT EXISTS credentials_id_seq;
CREATE TABLE IF NOT EXISTS credentials
(
id int8 NOT NULL DEFAULT nextval('credentials_id_seq'::regclass) PRIMARY KEY, 
user_id int8,
password varchar,
e int,
is_active bool default false,
created_at timestamptz,
updated_at timestamptz,
deleted_at timestamptz
);

-- +migrate Down
DROP SEQUENCE IF EXISTS credentials_seq;
DROP TABLE IF EXISTS credentials;
