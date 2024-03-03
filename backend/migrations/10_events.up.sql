BEGIN;

CREATE TABLE events
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    type       text NOT NULL,
    mark int NOT NULL,
    student_id uuid NOT NULL REFERENCES students (id),
    subject_id uuid NOT NULL REFERENCES subjects (id)
);

COMMIT;
