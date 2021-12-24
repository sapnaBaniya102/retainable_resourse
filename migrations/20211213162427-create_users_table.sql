
-- +migrate Up
CREATE SEQUENCE IF NOT EXISTS users_id_seq;
CREATE TABLE IF NOT EXISTS users
(
id int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass) PRIMARY KEY, 
email varchar(250),
user_id int8,
status varchar(50),
is_active bool default false,
created_at timestamptz,
updated_at timestamptz,
deleted_at timestamptz
);

-- +migrate Down
DROP SEQUENCE IF EXISTS users_seq;
DROP TABLE IF EXISTS users;
