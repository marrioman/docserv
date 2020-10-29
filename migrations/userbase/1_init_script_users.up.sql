CREATE SCHEMA IF NOT EXISTS users;
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS users.user_profile (
    id text DEFAULT gen_random_uuid(),
    user_name text,
    password varchar(20),
    created_at timestamptz not null default now(),
    last_name varchar,
    primary key (id)
);

INSERT INTO users.user_profile (id, user_name, password, last_name)
VALUES ('48ab8b9d-e14c-49fe-bdd7-f439e64bb6ae', 'killlo', 'qwerty', 'Бурдеев'),
       ('ddc721a1-5aaf-4fee-a548-f65e678c28b7', 'wowzer', 'qwerty', 'Иванов');