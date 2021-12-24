
-- +migrate Up
CREATE SEQUENCE IF NOT EXISTS user_profiles_id_seq;
CREATE TABLE IF NOT EXISTS user_profiles
(
id int8 NOT NULL DEFAULT nextval('user_profiles_id_seq'::regclass) PRIMARY KEY, 
firstname varchar(150),
lastname varchar(200),
username varchar(200),

status varchar(50),

is_active bool default false,
created_at timestamptz,
updated_at timestamptz,
deleted_at timestamptz
);

-- +migrate Down
DROP SEQUENCE IF EXISTS user_profiles_seq;
DROP TABLE IF EXISTS user_profiles;
