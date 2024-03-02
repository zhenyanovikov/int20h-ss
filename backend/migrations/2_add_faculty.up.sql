BEGIN;

create table faculties
(
    id   uuid default uuid_generate_v4(),
    name text
);

COMMIT;