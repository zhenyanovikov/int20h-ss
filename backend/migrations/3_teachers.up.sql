BEGIN;

CREATE TABLE teachers
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id    uuid NOT NULL REFERENCES users (id),
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

COMMIT;
