CREATE TABLE newsletters
(
    id           UUID PRIMARY KEY,
    title        TEXT        NOT NULL,
    content      TEXT        NOT NULL,
    attachment   TEXT,
    created_at   TIMESTAMPTZ NOT NULL,
    scheduled_at TIMESTAMPTZ
);