BEGIN;

CREATE TABLE notification_templates
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    subject    TEXT NOT NULL,
    body       TEXT NOT NULL,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

COMMIT;
