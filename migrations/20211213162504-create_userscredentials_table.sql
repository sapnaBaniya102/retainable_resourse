
-- +migrate Up
CREATE SEQUENCE IF NOT EXISTS userscredentials_id_seq;
CREATE TABLE IF NOT EXISTS userscredentials
(
id int8 NOT NULL DEFAULT nextval('userscredentials_id_seq'::regclass) PRIMARY KEY, 
user_id int8,
password varchar(250),
status varchar(50),
is_active bool default false,
created_at timestamptz,
updated_at timestamptz,
deleted_at timestamptz
);

-- +migrate Down
DROP SEQUENCE IF EXISTS userscredentials_seq;
DROP TABLE IF EXISTS userscredentials;
