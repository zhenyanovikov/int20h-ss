BEGIN;

create table faculties
(
    id   uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text
);

COMMIT;