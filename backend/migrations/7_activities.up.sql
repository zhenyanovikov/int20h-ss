BEGIN;

CREATE TABLE activities
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       TEXT NOT NULL,
    student_id uuid NOT NULL REFERENCES students (id) ON DELETE CASCADE,
    bonus_mark integer,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

COMMIT;
