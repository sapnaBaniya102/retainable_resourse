
-- +migrate Up
CREATE SEQUENCE IF NOT EXISTS profiles_id_seq;
CREATE TABLE IF NOT EXISTS profiles
(
id int8 NOT NULL DEFAULT nextval('profiles_id_seq'::regclass) PRIMARY KEY, 
firstname varchar(150),
lastname varchar(150),
username varchar(250),
email varchar(250),
status varchar(60),
is_active bool default false,
created_at timestamptz,
updated_at timestamptz,
deleted_at timestamptz
);

-- +migrate Down
DROP SEQUENCE IF EXISTS profiles_seq;
DROP TABLE IF EXISTS profiles;
