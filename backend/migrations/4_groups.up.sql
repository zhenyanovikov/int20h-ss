BEGIN;

CREATE TABLE groups (
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       text NOT NULL,
    year_start integer NOT NULL,
    faculty_id uuid NOT NULL REFERENCES faculties(id)
);

COMMIT;
