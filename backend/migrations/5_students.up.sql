BEGIN;

CREATE TABLE students
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id    uuid NOT NULL REFERENCES users (id),
    group_id   uuid REFERENCES groups (id),
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

COMMIT;
