BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users
(
    id         uuid          default uuid_generate_v4(),
    first_name text not null,
    last_name  text not null default '',
    email      text not null unique,
    avatar_url text,
    primary key (id)
);

create table user_credentials
(
    id           uuid not null default uuid_generate_v4(),
    user_id      uuid references users (id) on delete cascade,
    access_token text,
    expires_at   text,
    created_at   timestamp     default now(),
    primary key (id)
);

COMMIT;
