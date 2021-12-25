
-- +migrate Up
CREATE SEQUENCE IF NOT EXISTS login_sessions_id_seq;
CREATE TABLE IF NOT EXISTS login_sessions
(
id int8 NOT NULL DEFAULT nextval('login_sessions_id_seq'::regclass) PRIMARY KEY, 
  "user_id"  int8,
    "k"        varchar(64),
    "v"        bytea,
    "e"        int8 DEFAULT 0,
is_active bool default false,
created_at timestamptz,
updated_at timestamptz,
deleted_at timestamptz
);

-- +migrate Down
DROP SEQUENCE IF EXISTS login_sessions_seq;
DROP TABLE IF EXISTS login_sessions;
