BEGIN;

CREATE TABLE subjects
(
    id         uuid unique DEFAULT uuid_generate_v4(),
    teacher_id uuid NOT NULL REFERENCES teachers (id),
    group_id   uuid NOT NULL REFERENCES groups (id),
    name text NOT NULL,

    PRIMARY KEY (id, teacher_id, group_id, name)
);

COMMIT;
