BEGIN;

CREATE TABLE assignments
(
    id             uuid PRIMARY KEY   DEFAULT uuid_generate_v4(),
    subject_id     uuid      NOT NULL REFERENCES subjects (id),

    title          text      NOT NULL,
    description    text,
    attachment_url text,
    deadline       timestamp NOT NULL,

    created_at     timestamp NOT NULL DEFAULT now()
);

CREATE TABLE submitted_assignments
(
    id             uuid PRIMARY KEY   DEFAULT uuid_generate_v4(),
    assignment_id  uuid      NOT NULL REFERENCES assignments (id),
    student_id     uuid      NOT NULL REFERENCES students (id),

    text           text,
    attachment_url text,
    eventID        uuid,
    created_at     timestamp NOT NULL DEFAULT now()
);

COMMIT;
